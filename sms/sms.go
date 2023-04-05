package sms

import (
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	"github.com/chenyuIT/framework/facades"
)

type Sender interface {
	Send(sms *SMS) error
}

var SenderMap = make(map[string]func() Sender)

func sendcode(sms *SMS) error {
	vendor := sms.Config.Vendor
	if s, ok := SenderMap[vendor]; ok {
		return s().Send(sms)
	}
	//强制要求设置Config.Vendor这个参数
	panic("设置的短信服务商有误")
}

type ServiceConfig struct {
	Vendor      string
	Group       string
	Tpl         string
	Signname    string
	Allowcity   []string
	MaxSendNums int    //每天最大发送数量
	Callback    string //成功后回调URL
	Mode        byte
	Validtime   int
	Outformat   string
}

type SMS struct {
	Mobile      string
	Code        string
	Uid         string
	serviceName string
	Config      *ServiceConfig
	ConfigisOK  bool
	NowTime     time.Time
	model       *Model
}

func Init() {
	time.Local = func() *time.Location {
		loc, err := time.LoadLocation(facades.Config.GetString("sms.timezone"))
		if err != nil {
			loc = time.UTC
		}
		return loc
	}()
}

func NewSms() *SMS {
	sms := &SMS{}
	sms.Config = &ServiceConfig{}
	sms.model = NewModel(sms)
	sms.NowTime = time.Now()
	Init()
	sms.SetServiceConfig("register")

	SenderMap["alidayu"] = func() Sender {
		return &Alidayu{}
	}

	return sms
}

// 设置服务配置文件，根据register/resetpwd选择不同短信模板，暂时使用同一模板，
func (sms *SMS) SetServiceConfig(serviceName string) *SMS {
	configs := facades.Config.Get(fmt.Sprintf("sms.vendors.%s.servicelist.%s", facades.Config.GetString("sms.default"), serviceName))
	sms.fillDefaultForConfigs(configs.(map[string]any))
	sms.ConfigisOK = true
	if sms.ConfigisOK {
		sms.serviceName = serviceName
		return sms
	}
	//强制要求设置serviceName这个参数
	panic("服务" + serviceName + "配置不存在!")
}

func (sms *SMS) fillDefaultForConfigs(config map[string]any) {
	sms.Config.Vendor = config["vendor"].(string)
	sms.Config.Group = config["group"].(string)
	sms.Config.Tpl = config["smstpl"].(string)
	sms.Config.Signname = config["signname"].(string)
	// sms.Config.Allowcity   = []{"a","b"}
	sms.Config.MaxSendNums = config["maxsendnums"].(int)
	// sms.Config.Callback = ""
	sms.Config.Mode = 3
	sms.Config.Validtime = config["validtime"].(int)
	sms.Config.Outformat = config["outformat"].(string)
}

// 归属地规则校验
func (sms *SMS) checkArea() error {

	if len(sms.Config.Allowcity) < 1 { //没有启用
		return nil
	}

	area, err := sms.model.GetMobileArea()
	if err != nil {
		return err
	}

	var Allow = false
	for _, citycode := range sms.Config.Allowcity {
		if strings.Contains(area, citycode) {
			Allow = true //允许发送sms
			break
		}
	}

	if !Allow {
		errorMsg := facades.Config.GetString("sms.errormsg.err_allow_areacode")
		return fmt.Errorf(errorMsg, strings.Join(sms.Config.Allowcity, ","))
	}

	return nil
}

func (sms *SMS) checkhold() error {

	sendTime, err := sms.model.GetSendTime()
	if err != nil {
		return err
	}

	if sendTime > 0 && sms.NowTime.Unix()-sendTime < Maxsendtime { //发送间隔不能小于60秒
		errorMsg := facades.Config.GetString("sms.errormsg.err_per_minute_send_num")
		return fmt.Errorf(errorMsg)
	}

	sendMax, err := sms.model.GetTodaySendNums()
	if err != nil {
		return err
	}

	if sendMax > 0 && sendMax >= (uint64)(sms.Config.MaxSendNums) {
		errorMsg := facades.Config.GetString("sms.errormsg.err_per_day_max_send_nums")
		return fmt.Errorf(errorMsg, sms.Config.MaxSendNums)
	}

	return nil
}

/*
*
当前模式  1：只有手机号对应的uid存在时才能发送，2：只有uid不存在时才能发送，3：不管uid是否存在都发送
*
*/
func (sms *SMS) currModeok() error {

	uid, err := sms.model.GetSmsUid()
	if err != nil {
		return err
	}
	switch mode := sms.Config.Mode; mode {
	case 0x01:
		if uid != "" {
			return nil
		}
		errorMsg := facades.Config.GetString("sms.errormsg.err_model_not_ok1")
		return fmt.Errorf(errorMsg, sms.Mobile)
	case 0x02:
		if uid == "" {
			return nil
		}
		errorMsg := facades.Config.GetString("sms.errormsg.err_model_not_ok1")
		return fmt.Errorf(errorMsg, sms.Mobile)
	case 0x03:
		return nil
	}
	//强制要求设置这个模式参数，有利于更加清楚服务调用者明确发送验证码与uid之间的关联
	panic(fmt.Errorf("请正确配置对应服务中的mode参数"))
}

// 保存数据
func (sms *SMS) save() {

	sms.model.SetSendTime()

	nums, _ := sms.model.GetTodaySendNums()

	newnums := atomic.AddUint64(&nums, 1) //原子操作+1

	sms.model.SetTodaySendNums(newnums)

	sms.model.SetSmsCode()
}

// 发送短信
// 需保证在高并发下同一个手机号相同短信服务send操作是同步的，确保后续规则校验可以依次进行；
func (sms *SMS) Send(mobile string) error {
	if !sms.ConfigisOK {
		//强制要求明确服务参数配置
		panic(fmt.Errorf("(%s)服务配置不存在", sms.serviceName))
	}

	if err := VailMobile(mobile); err != nil {
		return err
	}

	sms.Mobile = mobile
	sms.Code = makeCode()

	if err := sms.checkArea(); err != nil {
		return err
	}
	if err := sms.currModeok(); err != nil {
		return err
	}
	if err := sms.checkhold(); err != nil {
		return err
	}
	if err := sendcode(sms); err != nil {

		//发送失败 callback
		AddCallbackTask(sms, "Failed")
		return err
	}

	//保存记录
	sms.save()

	//发送成功 callback
	AddCallbackTask(sms, "Success")

	return nil
}

func (sms *SMS) CheckCode(mobile, code string) error {
	if !sms.ConfigisOK {
		panic(fmt.Errorf("(%s)服务配置不存在", sms.serviceName))
	}

	sms.Mobile = mobile
	sms.Code = code

	if err := VailMobile(sms.Mobile); err != nil {
		return err
	}

	if err := VailCode(sms.Code); err != nil {
		return err
	}

	oldcode, validtime, _ := sms.model.GetSmsCode()

	if oldcode == "" || sms.Code != oldcode {
		errorMsg := facades.Config.GetString("sms.errormsg.err_code_not_ok")
		return fmt.Errorf(errorMsg, sms.Code)
	}

	if sms.NowTime.Unix() > validtime {
		time1 := time.Unix(validtime, 0)
		errorMsg := facades.Config.GetString("sms.errormsg.err_vailtime_not_ok")
		return fmt.Errorf(errorMsg, time.Since(time1).String())
	}

	//验证成功时 callback
	AddCallbackTask(sms, "Checkok")

	return nil
}

func (sms *SMS) SetUid(mobile, uid string) error {
	if !sms.ConfigisOK {
		panic(fmt.Errorf("(%s)服务配置不存在", sms.serviceName))
	}

	sms.Mobile = mobile
	sms.Uid = uid

	if err := VailMobile(sms.Mobile); err != nil {
		return err
	}

	if err := VailUid(sms.Uid); err != nil {
		return err
	}

	sms.model.SetSmsUid()

	return nil
}

func (sms *SMS) DelUid(mobile, uid string) error {
	if !sms.ConfigisOK {
		panic(fmt.Errorf("(%s)服务配置不存在", sms.serviceName))
	}

	sms.Mobile = mobile
	sms.Uid = uid

	if err := VailMobile(sms.Mobile); err != nil {
		return err
	}
	if err := VailUid(sms.Uid); err != nil {
		return err
	}

	olduid, err := sms.model.GetSmsUid()

	errorMsg := facades.Config.GetString("sms.errormsg.err_not_uid")
	if err != nil {
		return fmt.Errorf(errorMsg, sms.Mobile, sms.Uid)
	}
	if olduid != uid {
		return fmt.Errorf(errorMsg, sms.Mobile, sms.Uid)
	}

	sms.model.DelSmsUid()
	return nil
}

func (sms *SMS) Info(mobile string) (map[string]interface{}, error) {
	if !sms.ConfigisOK {
		panic(fmt.Errorf("(%s)服务配置不存在", sms.serviceName))
	}

	sms.Mobile = mobile

	if err := VailMobile(sms.Mobile); err != nil {
		return nil, err
	}
	info := make(map[string]interface{})
	info["mobile"] = sms.Mobile
	info["service"] = sms.serviceName
	info["areacode"], _ = sms.model.GetMobileArea()
	info["lastsendtime"], _ = sms.model.GetSendTime()
	info["sendnums"], _ = sms.model.GetTodaySendNums()
	info["smscode"], info["smscodeinvalidtime"], _ = sms.model.GetSmsCode()
	info["extinfo"], _ = sms.model.GetMobileInfo()
	info["uid"], _ = sms.model.GetSmsUid()
	return info, nil
}
