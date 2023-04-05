package ocr

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/chenyuIT/framework/contracts/ocr"
	"github.com/chenyuIT/framework/facades"
)

var token string

func Init() {
	ok, data := GetToken()
	fmt.Println(ok, data)
	if ok != nil {
		fmt.Println(ok)
		return
	}
	token = data.Token
}

type ResponseData struct {
	Code  string
	Msg   string
	Token string
}

type OCR struct {
}

func NewOcr() *OCR {
	ocr := &OCR{}
	Init()
	return ocr
}

// OcrImg 调用百度文字识别api
// func (ocr *OCR) OcrImg(base string) (data ReturnData, err error) {
func (_ocr *OCR) OcrImg(base string) (data ocr.ReturnData, err error) {
	if !strings.HasPrefix(base, "data:image") {
		var dt ocr.ReturnData
		return dt, errors.New("param format error")
	}
	err, data = YunOcr(token, base)
	return
}

// GetIOcrImg 调用百度文字识别自定义api
func (_ocr *OCR) GetIOcrImg(templateId string, base string) (data ocr.ReturnIOcrData, err error) {
	err, data = iOcr(token, templateId, base)
	return
}

func GetToken() (error, ResponseData) {
	response := ResponseData{}
	//读取token文件
	file, err := os.OpenFile("./ocr.tk", os.O_RDWR, 0600)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		file, createrr := os.Create("./ocr.tk")
		if createrr != nil {
			fmt.Println("创建token文件失败")
			return createrr, response
		}
		rterr, response := authtoken()
		if rterr == nil {
			_, err = file.Write([]byte(response.Token))
		}
		return rterr, response
	}

	//校验token文件的有效期
	fileinfo, _ := file.Stat()
	changetime := fileinfo.ModTime()
	hours := time.Now().Sub(changetime).Hours()
	fmt.Println(hours)
	//大于29天就重新获取 实际有效期是30天
	if int(hours) >= 29*24 {
		rterr, response := authtoken()
		if rterr == nil {
			_, err = file.Write([]byte(response.Token))
			if err != nil {
				fmt.Println(err)
				return err, response
			}
		}
		return rterr, response
	}
	buff := make([]byte, 1000)
	n, err := file.Read(buff)
	fmt.Println(string(buff[:n]))
	if err != nil {
		return nil, response
	}
	response.Token = string(buff[:n])
	return err, response
}

func authtoken() (error, ResponseData) {
	tokenUrl := facades.Config.GetString(fmt.Sprintf("ocr.vendors.%s.tokenUrl", facades.Config.GetString("ocr.default")))
	granttype := facades.Config.GetString(fmt.Sprintf("ocr.vendors.%s.granttype", facades.Config.GetString("ocr.default")))
	clientid := facades.Config.GetString(fmt.Sprintf("ocr.vendors.%s.clientid", facades.Config.GetString("ocr.default")))
	clientsecret := facades.Config.GetString(fmt.Sprintf("ocr.vendors.%s.clientsecret", facades.Config.GetString("ocr.default")))

	data := url.Values{"grant_type": {granttype}, "client_id": {clientid}, "client_secret": {clientsecret}}
	respmsg, msgerr := http.PostForm(tokenUrl, data)
	// fmt.Println(respmsg.Request.Form, respmsg.Request.Method, respmsg.Request.URL)
	response := ResponseData{}
	if msgerr != nil {
		response.Code = "400"
		response.Msg = "token请求失败"
		response.Token = ""
		return msgerr, response
	}
	defer respmsg.Body.Close()
	resbody, reserr := ioutil.ReadAll(respmsg.Body)
	if reserr != nil {
		// handle error
		response.Code = "400"
		response.Msg = "token读取失败"
		response.Token = ""
		return reserr, response
	}
	var rtdata struct {
		Access_token string
		Expires_in   int
	}
	err := json.Unmarshal(resbody, &rtdata)
	if err != nil {
		fmt.Println(err)
		response.Code = "400"
		response.Msg = "token解析失败"
		response.Token = ""
		return err, response
	}
	response.Code = ""
	response.Msg = "token读取成功"
	response.Token = rtdata.Access_token
	return nil, response
}

func YunOcr(token, baseStr string) (error, ocr.ReturnData) {
	ocrAccurateUrl := facades.Config.GetString(fmt.Sprintf("ocr.vendors.%s.ocrAccurateUrl", facades.Config.GetString("ocr.default")))
	sendUrl := ocrAccurateUrl + "?access_token=" + token
	data := url.Values{"image": {baseStr}}
	var client = http.DefaultClient
	req, err := http.NewRequest("POST", sendUrl, strings.NewReader(data.Encode()))

	var returnData ocr.ReturnData
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rp, err := client.Do(req)
	if err != nil {
		return err, returnData
	}
	defer rp.Body.Close()
	resBody, resErr := ioutil.ReadAll(rp.Body)
	if resErr != nil {
		return resErr, returnData
	}

	err = json.Unmarshal(resBody, &returnData)
	if err != nil {
		fmt.Println("无法解析")
	}
	return err, returnData
}

func iOcr(token, templateId string, baseStr string) (error, ocr.ReturnIOcrData) {
	iOcrUrl := facades.Config.GetString(fmt.Sprintf("ocr.vendors.%s.iOcrUrl", facades.Config.GetString("ocr.default")))
	sendUrl, err := url.Parse(iOcrUrl)
	if err != nil {
		fmt.Println(err)
	}
	query := sendUrl.Query()
	query.Set("access_token", token)
	sendUrl.RawQuery = query.Encode()

	sendBody := http.Request{}
	sendBody.ParseForm()
	sendBody.Form.Add("image", baseStr)
	sendBody.Form.Add("templateSign", templateId)
	sendData := sendBody.Form.Encode()
	var client = http.DefaultClient
	req, err := http.NewRequest("POST", sendUrl.String(), strings.NewReader(sendData))

	var returnData ocr.ReturnIOcrData
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	result, err := client.Do(req)
	if err != nil {
		return err, returnData
	}
	defer result.Body.Close()
	resBody, resErr := ioutil.ReadAll(result.Body)
	if resErr != nil {
		return resErr, returnData
	}
	err = json.Unmarshal(resBody, &returnData)

	if err != nil {
		return err, returnData
	}
	return err, returnData
}
