package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

var GetTokenTime int64 = 0
var Token string

type GetUserInfoStruct struct {
	Openid string `json:"openid"`
}

// JsonResult 返回结构
type JsonResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}

func ReturnBack(w http.ResponseWriter, r *http.Request, res JsonResult) {
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误010")
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}

// IndexText 计数器接口
func IndexText(w http.ResponseWriter, r *http.Request) {
	data, err := getFile("./MP_verify_X0kqrTo5XxsuQ4bB.txt")
	if err != nil {
		fmt.Fprint(w, "内部错误00:"+err.Error())
		return
	}
	fmt.Fprint(w, data)
}

// IndexHandler 计数器接口
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data, err := getFile("./index.html")
	if err != nil {
		fmt.Fprint(w, "内部错误01")
		return
	}
	fmt.Fprint(w, data)
}

// Index2Handler 计数器接口
func Index2Handler(w http.ResponseWriter, r *http.Request) {
	data, err := getFile("./index2.html")
	if err != nil {
		fmt.Fprint(w, "内部错误02")
		return
	}
	fmt.Fprint(w, data)
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0

	BodyBytes, _ := ioutil.ReadAll(r.Body)
	j := GetUserInfoStruct{}
	err := json.Unmarshal(BodyBytes, &j)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = "消息错误"
		ReturnBack(w, r, *res)
		return
	}
	Openid := r.Header.Get("X-Wx-Openid")
	at := r.Header.Get("X-Wx-Cloudbase-Access-Token")
	url := "https://api.weixin.qq.com/wxa/getopendata?openid=" + Openid + "&cloudbase_access_token=" + at + "&cloudid_list=" + j.Openid
	data := "{\"cloudid_list\": [\"" + j.Openid + "\"]}"
	payload := strings.NewReader(data)
	response, _ := http.Post(url, "application/json", payload)
	res.Code = 0
	res.ErrorMsg = ""
	BodyBytes0, _ := ioutil.ReadAll(response.Body)
	w.Header().Set("content-type", "application/json")
	w.Write(BodyBytes0)
}

func CheckAdmin(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	BodyBytes, _ := ioutil.ReadAll(r.Body)
	j := GetUserInfoStruct{}
	err := json.Unmarshal(BodyBytes, &j)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = "消息错误"
		ReturnBack(w, r, *res)
		return
	}
	res.Code = 0
	res.ErrorMsg = ""
	res.Data = dao.Imp.CheckAdminLevel(j.Openid)
	ReturnBack(w, r, *res)
	return
}
func UserLogin(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	BodyBytes, _ := ioutil.ReadAll(r.Body)
	j := &model.UserLoginStruct{}
	err := json.Unmarshal(BodyBytes, &j)
	if err != nil {
		res.Code = -1
		res.ErrorMsg = "消息错误"
		ReturnBack(w, r, *res)
		return
	}
	res.Code = 0
	res.ErrorMsg = ""
	res.Data = dao.Imp.UserLogin(j.Username, j.Pwd)
	ReturnBack(w, r, *res)
	return
}

// CarOver 计数器接口
func CarOver(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	if r.Method == http.MethodPost {
		BodyBytes, _ := ioutil.ReadAll(r.Body)
		counter := &model.WeihuapinCarOver{}
		err := json.Unmarshal(BodyBytes, &counter)
		if err != nil {
			res.Code = -2
			res.ErrorMsg = "消息结构体错误"
			ReturnBack(w, r, *res)
			return
		}
		counter.Outtime = time.Now()
		counter.Status = 1
		counter.Outtype = 0
		err = dao.Imp.OverCar(counter)

	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	ReturnBack(w, r, *res)
	return
}
func CarInsert(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	if r.Method == http.MethodPost {
		BodyBytes, _ := ioutil.ReadAll(r.Body)
		counter := &model.WeihuapinCarInsert{}
		err := json.Unmarshal(BodyBytes, &counter)
		if err != nil {
			res.Code = -2
			res.ErrorMsg = "消息结构体错误"
			ReturnBack(w, r, *res)
			return
		}
		counter.Intime = time.Now()
		counter.Recordtime = time.Now()
		counter.Outtype = 0
		counter.Status = 0
		err = dao.Imp.InsertCar(counter)

	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	ReturnBack(w, r, *res)
	return
}

// CarUpdate 计数器接口
func CarUpdate(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	if r.Method == http.MethodPost {
		BodyBytes, _ := ioutil.ReadAll(r.Body)
		counter := &model.WeihuapinCarUpdate{}
		err := json.Unmarshal(BodyBytes, &counter)
		if err != nil {
			res.Code = -2
			res.ErrorMsg = "消息结构体错误"
			ReturnBack(w, r, *res)
			return
		}
		if counter.Id <= 0 {
			res.Code = -21
			res.ErrorMsg = "id错误"
			ReturnBack(w, r, *res)
			return
		}
		err = dao.Imp.UpdateCar(counter)
	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	ReturnBack(w, r, *res)
	return
}

// CarGet 计数器接口
func CarGet(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	if r.Method == http.MethodPost {
		BodyBytes, _ := ioutil.ReadAll(r.Body)
		counter := &model.WeihuapinGetOne{}
		err := json.Unmarshal(BodyBytes, &counter)
		if err != nil || counter.WechartID == "" {
			res.Code = -3
			res.ErrorMsg = "消息结构体错误"
			ReturnBack(w, r, *res)
			return
		}

		res.Data, res.ErrorMsg, res.Code = dao.Imp.GetCar(counter.WechartID)

	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	ReturnBack(w, r, *res)
	return
}

// CarList 计数器接口
func CarList(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	if r.Method == http.MethodPost {
		BodyBytes, _ := ioutil.ReadAll(r.Body)
		counter := &model.WeihuapinGetList{}
		err := json.Unmarshal(BodyBytes, &counter)
		if err != nil {
			res.Code = -4
			res.ErrorMsg = "消息结构体错误"
			ReturnBack(w, r, *res)
			return
		}
		if counter.Status != 0 && counter.Status != 1 {
			counter.Status = 100
		}
		res.Data, res.ErrorMsg, res.Code = dao.Imp.GetRecord(counter.Status, counter.Offset, counter.Limit)

	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	ReturnBack(w, r, *res)
	return
}

// CarListNum 计数器接口
func CarListNum(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	if r.Method == http.MethodPost {
		BodyBytes, _ := ioutil.ReadAll(r.Body)
		counter := &model.WeihuapinGetList{}
		err := json.Unmarshal(BodyBytes, &counter)
		if err != nil {
			res.Code = -4
			res.ErrorMsg = "消息结构体错误"
			ReturnBack(w, r, *res)
			return
		}
		if counter.Status != 0 && counter.Status != 1 {
			counter.Status = 100
		}
		res.Data, res.ErrorMsg, res.Code = dao.Imp.GetRecordNum(counter.Status, counter.Offset, counter.Limit)

	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	ReturnBack(w, r, *res)
	return
}

// GetWeihuapin 计数器接口
func GetWeihuapin(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	if r.Method == http.MethodPost {
		BodyBytes, _ := ioutil.ReadAll(r.Body)
		counter := &model.WeihuapinInfo{}
		err := json.Unmarshal(BodyBytes, &counter)
		if err != nil {
			res.Code = -4
			res.ErrorMsg = "消息结构体错误"
			ReturnBack(w, r, *res)
			return
		}

		res.Data, res.ErrorMsg, res.Code = dao.Imp.GetWeihuapin(counter.Weihuapin)

	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	ReturnBack(w, r, *res)
	return
}

func AdminAdd(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	if r.Method == http.MethodPost {
		BodyBytes, _ := ioutil.ReadAll(r.Body)
		counter := &model.AdminInsert{}
		err := json.Unmarshal(BodyBytes, &counter)
		if err != nil {
			res.Code = -4
			res.ErrorMsg = "消息结构体错误"
		}
		counter.CreateTime = time.Now()
		if counter.Level <= 0 {
			counter.Level = 1
		}
		err2 := dao.Imp.InsertAdmin(counter)
		if err2 != nil {
			res.ErrorMsg, res.Code = err.Error(), -1
		}

	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}
	ReturnBack(w, r, *res)
	return
}

func AdminUpdate(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	if r.Method == http.MethodPost {
		BodyBytes, _ := ioutil.ReadAll(r.Body)
		counter := &model.AdminUpdate{}
		err := json.Unmarshal(BodyBytes, &counter)
		if err != nil {
			res.Code = -4
			res.ErrorMsg = "消息结构体错误"
			ReturnBack(w, r, *res)
			return
		}
		if counter.ID <= 0 {
			res.Code = -5
			res.ErrorMsg = "更新数据时，缺少ID"
			ReturnBack(w, r, *res)
			return
		}

		err2 := dao.Imp.UpdateAdmin(counter)
		if err2 != nil {
			res.ErrorMsg, res.Code = err.Error(), -1
		}

	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}
	ReturnBack(w, r, *res)
	return
}

func AdminOver(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	if r.Method == http.MethodPost {
		BodyBytes, _ := ioutil.ReadAll(r.Body)
		counter := &model.AdminOver{}
		err := json.Unmarshal(BodyBytes, &counter)
		if err != nil {
			res.Code = -4
			res.ErrorMsg = "消息结构体错误"
			ReturnBack(w, r, *res)
			return
		}
		if counter.ID <= 0 {
			res.Code = -5
			res.ErrorMsg = "删除数据时，缺少ID"
			ReturnBack(w, r, *res)
			return
		}

		err2 := dao.Imp.OverAdmin(counter)
		if err2 != nil {
			res.ErrorMsg, res.Code = err.Error(), -1
		}

	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}
	ReturnBack(w, r, *res)
	return
}

// AdminList 计数器接口
func AdminList(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	if r.Method == http.MethodPost {
		res.Data, res.ErrorMsg, res.Code = dao.Imp.GetAdminList()
	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	ReturnBack(w, r, *res)
	return
}

// getIndex 获取主页
func getFile(FileName string) (string, error) {
	b, err := ioutil.ReadFile(FileName)
	if err != nil {
		fmt.Println("读取文件出错", err)
		return "", err
	}
	return string(b), nil
}

func GetToken() {
	tnow := time.Now().Unix()
	if tnow-GetTokenTime < 7200 {
		return
	}
	GetTokenTime = tnow
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxa806018a131603d3&secret=2e23f64187fa7aed8d528c5d0451288a"
	response, _ := http.Post(url, "application/json", nil)
	BodyBytes0, _ := ioutil.ReadAll(response.Body)
	fmt.Println("GetToken==", string(BodyBytes0))

	//url := "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send？cloudbase_access_token=ff"
	//data := "{\"touser\":\"oCZvY55E1d6jDmo5wnGhvbZ4ROoo\",\"template_id\":\"o2yh8P7T9K3rR9f4H_DxGpwYQntM4b3ZCv3EUnfwTQs\",\"url\":\"\",\"topcolor\":\"#FF0000\",\"data\":{\"first\":{\"value\":\"尊敬的 京A00001 车主，您已停车两个小时：\",\"color\":\"#173177\"},\"keyword1\":{\"value\":\"济南东高速服务区危化品车辆停车场\",\"color\":\"#173177\"},\"keyword2\":{\"value\":\"2022-09-06 08:49:00\",\"color\":\"#173177\"},\"keyword3\":{\"value\":\"--\",\"color\":\"#173177\"},\"keyword4\":{\"value\":\"--\",\"color\":\"#173177\"},\"keyword5\":{\"value\":\"--\",\"color\":\"#173177\"},\"remark\":{\"value\":\"祝您出行愉快！\",\"color\":\"#173177\"}}}"
	//payload := strings.NewReader(data)
	//response, _ := http.Post(url, "application/json", payload)
	//
	//BodyBytes0, _ := ioutil.ReadAll(response.Body)
	//
	//fmt.Println(string(BodyBytes0))
}
func SendMsg(msg map[string]interface{}) {
	templateid := "o2yh8P7T9K3rR9f4H_DxGpwYQntM4b3ZCv3EUnfwTQs"
	url := "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send？cloudbase_access_token=" + Token
	data := "{\"touser\":\"" + msg["wechartid"].(string) + "\",\"mp_template_msg\":{\"appid\":\"wxa806018a131603d3\",\"template_id\":\"" + templateid + "\",\"url\":\"\",\"topcolor\":\"#FF0000\",\"miniprogram\":{\"appid\":\"wx032125bd60fe9474\",\"pagepath\":\"\"},\"data\":{\"first\":{\"value\":\"尊敬的 " + msg["CarNo"].(string) + " 车主，您已停车两个小时：\",\"color\":\"#173177\"},\"keyword1\":{\"value\":\"济南东高速服务区危化品车辆停车场\",\"color\":\"#173177\"},\"keyword2\":{\"value\":\"" + msg["intime"].(string) + "\",\"color\":\"#173177\"},\"keyword3\":{\"value\":\"--\",\"color\":\"#173177\"},\"keyword4\":{\"value\":\"--\",\"color\":\"#173177\"},\"keyword5\":{\"value\":\"--\",\"color\":\"#173177\"},\"remark\":{\"value\":\"祝您出行愉快！\",\"color\":\"#173177\"}}}}"
	payload := strings.NewReader(data)
	response, _ := http.Post(url, "application/json", payload)
	//
	BodyBytes0, _ := ioutil.ReadAll(response.Body)
	//
	fmt.Println("发送消息结果", string(BodyBytes0))
}

func ScanData() {

	for {
		GetToken()
		data, _, _ := dao.Imp.GetMsg(7200)
		for _, v := range data {
			SendMsg(v)
		}
		time.Sleep(60 * 1000 * 1000 * 1000)
	}

}
