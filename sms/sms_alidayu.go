package sms

import (
	"github.com/chenyuIT/framework/facades"

	"github.com/opensource-conet/alidayu"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Alidayu struct {
	sms *SMS
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func (a *Alidayu) Send(sms *SMS) error {
	a.sms = sms
	alidayu.Appkey = facades.Config.GetString("sms.vendors.alidayu.appkey")
	alidayu.AppSecret = facades.Config.GetString("sms.vendors.alidayu.appSecret")

	client, _err := CreateClient(tea.String(alidayu.Appkey), tea.String(alidayu.AppSecret))
	if _err != nil {
		return _err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("阿里云短信测试"),
		TemplateCode:  tea.String("SMS_154950909"),
		PhoneNumbers:  tea.String("15818736970"),
		TemplateParam: tea.String("{\"code\":\"9856\"}"),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := client.SendSmsWithOptions(sendSmsRequest, runtime)
	if _err != nil {
		return _err
	}
	print(resp)
	// if !resp.Success {
	// 	log.V(1).Infof("Alidayu:%+v", res.ResultError)
	// 	return fmt.Errorf("%s", res.ResultError.SubMsg)
	// }
	return nil
}
