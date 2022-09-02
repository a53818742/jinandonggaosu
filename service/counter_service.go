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

func AdminAdd(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
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
