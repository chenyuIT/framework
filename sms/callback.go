package sms

import (
	"fmt"
	"net/url"
)

type Callback struct {
	Url      string
	Data     url.Values
	callnums uint
}

var (
	callback = make(chan *Callback)
)

func RunCallbackTask() {

}

func AddCallbackTask(sms *SMS, flag string) {

	if len(sms.Config.Callback) < 1 { //没有启用
		return
	}
	data := make(url.Values)
	data.Set("mobile", sms.Mobile)
	data.Set("code", sms.Code)
	data.Set("service", sms.serviceName)
	data.Set("uxtime", fmt.Sprintf("%d", sms.NowTime.Unix()))
	data.Set("flag", flag)
	callback <- &Callback{sms.Config.Callback, data, 0}
}

func (c *Callback) Do(cbs <-chan *Callback) {

}
