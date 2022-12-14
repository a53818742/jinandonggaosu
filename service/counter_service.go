package service

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

var GetTokenTime int64 = 0
var Token string

type TokenStruct struct {
	AccessToken string `json:"access_token"`
}
type SendMsgCallBack struct {
	ErrorCode int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}
type GetUserInfoStruct struct {
	Openid string `json:"openid"`
}
type RecordListStruct struct {
	UserID int `json:"userid"`
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

// Index3Handler 计数器接口
func Index3Handler(w http.ResponseWriter, r *http.Request) {
	data, err := getFile("./index3.html")
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
	res.Data = Openid

	ReturnBack(w, r, *res)

	//at := r.Header.Get("X-Wx-Cloudbase-Access-Token")
	//url := "https://api.weixin.qq.com/wxa/getopendata?openid=" + Openid + "&cloudbase_access_token=" + at + "&cloudid_list=" + j.Openid
	//data := "{\"cloudid_list\": [\"" + j.Openid + "\"]}"
	//payload := strings.NewReader(data)
	//response, e0 := http.Post(url, "application/json", payload)
	//if e0 == nil {
	//	res.Code = 0
	//	res.ErrorMsg = ""
	//
	//	BodyBytes0, e1 := ioutil.ReadAll(response.Body)
	//	BodyBytes0 := []byte("")
	//	if e1 == nil {
	//		w.Header().Set("content-type", "application/json")
	//		w.Write(BodyBytes0)
	//	} else {
	//		fmt.Println(e1)
	//	}
	//
	//} else {
	//	fmt.Println(e0)
	//}

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
func UserLogin2(w http.ResponseWriter, r *http.Request) {
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
	res.Data = dao.Imp.UserLogin2(j.Username, j.Pwd, j.Wechartid)
	ReturnBack(w, r, *res)
	return
}

func UserLogin3(w http.ResponseWriter, r *http.Request) {
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
	res.Data = dao.Imp.UserLogin3(j.Username, j.Pwd, j.Wechartid)
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
		res.Code, res.ErrorMsg = dao.Imp.InsertCar(counter)

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
		res.Data, res.ErrorMsg, res.Code = dao.Imp.GetRecordNum(counter.Status)

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

		//fmt.Println("GetWeihuapin", string(BodyBytes))
		counter := &model.WeihuapinInfo{}
		err := json.Unmarshal(BodyBytes, &counter)
		if err != nil {
			res.Code = -4
			res.ErrorMsg = "消息结构体错误"

			//fmt.Println("获取危化品", err)
			ReturnBack(w, r, *res)
			return
		}
		//fmt.Println("获取危化品", counter.Weihuapin)
		res.Data, res.ErrorMsg, res.Code = dao.Imp.GetWeihuapinByCN(counter.Weihuapin)

	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	ReturnBack(w, r, *res)
	return
}

// GetWeihuapinUN 计数器接口
func GetWeihuapinUN(w http.ResponseWriter, r *http.Request) {
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

		res.Data, res.ErrorMsg, res.Code = dao.Imp.GetWeihuapinByUN(counter.Weihuapin)

	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	ReturnBack(w, r, *res)
	return
}

func UserInsert(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	if r.Method == http.MethodPost {
		BodyBytes, _ := ioutil.ReadAll(r.Body)
		counter := &model.UserInsert{}
		err := json.Unmarshal(BodyBytes, &counter)
		if err != nil {
			res.Code = -4
			res.ErrorMsg = "消息结构体错误"
		}
		if counter.Level <= 0 {
			counter.Level = 1
		}
		res.Code, res.ErrorMsg = dao.Imp.UserAdd(counter)

	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}
	ReturnBack(w, r, *res)
	return
}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	if r.Method == http.MethodPost {
		BodyBytes, _ := ioutil.ReadAll(r.Body)
		counter := &model.UserUpdate{}
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

		err2 := dao.Imp.UserUpdate(counter)
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

func UserDelete(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	if r.Method == http.MethodPost {
		BodyBytes, _ := ioutil.ReadAll(r.Body)
		counter := &model.UserDelete{}
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

		err2 := dao.Imp.UserDelete(counter)
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

// UserList 计数器接口
func UserList(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	if r.Method == http.MethodPost {
		res.Data, res.ErrorMsg, res.Code = dao.Imp.UserList()
	} else {
		res.Code = -1
		res.ErrorMsg = fmt.Sprintf("请求方法 %s 不支持", r.Method)
	}

	ReturnBack(w, r, *res)
	return
}

func RecordInsert(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	if r.Method == http.MethodPost {
		BodyBytes, _ := ioutil.ReadAll(r.Body)
		counter := &model.RecordInsert{}
		err := json.Unmarshal(BodyBytes, &counter)
		if err != nil {
			res.Code = -4
			res.ErrorMsg = "消息结构体错误"
		}

		err2 := dao.Imp.RecordAdd(counter)
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

func RecordUpdate(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	if r.Method == http.MethodPost {
		BodyBytes, _ := ioutil.ReadAll(r.Body)
		counter := &model.RecordUpdate{}
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

		err2 := dao.Imp.RecordUpdate(counter)
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

func RecordDelete(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	if r.Method == http.MethodPost {
		BodyBytes, _ := ioutil.ReadAll(r.Body)
		counter := &model.RecordDelete{}
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

		err2 := dao.Imp.RecordDelete(counter)
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

// RecordList 计数器接口
func RecordList(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	res.Data = 0
	res.ErrorMsg = ""
	res.Code = 0
	if r.Method == http.MethodPost {

		BodyBytes, _ := ioutil.ReadAll(r.Body)
		counter := &RecordListStruct{}
		err := json.Unmarshal(BodyBytes, &counter)
		if err != nil {
			res.Code = -4
			res.ErrorMsg = "消息结构体错误"
			ReturnBack(w, r, *res)
			return
		}

		res.Data, res.ErrorMsg, res.Code = dao.Imp.RecordList(counter.UserID)
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
	if tnow-GetTokenTime < 4000 {
		return
	}

	url := "http://47.104.175.23:81/wechat_token.php?appid=wxa806018a131603d3&secret=2e23f64187fa7aed8d528c5d0451288a"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, er0 := client.Post(url, "application/json", nil)

	if er0 == nil {
		BodyBytes0, e1 := ioutil.ReadAll(response.Body)
		if e1 == nil {
			fmt.Println("GetToken ==", string(BodyBytes0))
			var msgstruct TokenStruct
			err11 := json.Unmarshal(BodyBytes0, &msgstruct)
			if err11 != nil {
				fmt.Println("GetToken Error11", err11)
				return
			}
			fmt.Println("msgstruct.AccessToken", msgstruct.AccessToken)
			if msgstruct.AccessToken != "" {
				GetTokenTime = tnow
				Token = msgstruct.AccessToken
			}
		} else {
			fmt.Println("GetToken Error1", e1)
		}
	} else {
		fmt.Println("GetToken Error", er0)
	}
}
func Interface2Int(inte interface{}) int {
	if reflect.TypeOf(inte).Kind() == reflect.String {
		d, _ := strconv.Atoi(inte.(string))
		return d
	}

	if reflect.TypeOf(inte).Kind() == reflect.Uint8 {
		return int(inte.(uint8))
	}
	if reflect.TypeOf(inte).Kind() == reflect.Int8 {
		return int(inte.(int8))
	}
	if reflect.TypeOf(inte).Kind() == reflect.Int16 {
		return int(inte.(int16))
	}
	if reflect.TypeOf(inte).Kind() == reflect.Int32 {
		return int(inte.(int32))
	}
	if reflect.TypeOf(inte).Kind() == reflect.Int {
		return inte.(int)
	}
	if reflect.TypeOf(inte).Kind() == reflect.Int64 {
		return int(inte.(int64))
	}
	if reflect.TypeOf(inte).Kind() == reflect.Float64 {
		return int(inte.(float64))
	}
	if reflect.TypeOf(inte).Kind() == reflect.Float32 {
		return int(inte.(float32))
	}
	return 0
}
func SendMsgToAdmin(msg map[string]interface{}, chartarray []string) {
	templateid := "o2yh8P7T9K3rR9f4H_DxGpwYQntM4b3ZCv3EUnfwTQs"
	url := "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + Token

	for _, wechartied := range chartarray {
		data := "{\"touser\":\"" + wechartied + "\",\"template_id\":\"" + templateid + "\",\"data\":{\"first\":{\"value\":\"" + msg["CarNo"].(string) + " 已停车超过两个小时，请及时处理！\",\"color\":\"#173177\"},\"keyword1\":{\"value\":\"济南东高速服务区危化品车辆停车场\",\"color\":\"#173177\"},\"keyword2\":{\"value\":\"" + msg["intime"].(time.Time).String()[0:19] + "\",\"color\":\"#173177\"},\"keyword3\":{\"value\":\"--\",\"color\":\"#173177\"},\"keyword4\":{\"value\":\"--\",\"color\":\"#173177\"},\"keyword5\":{\"value\":\"--\",\"color\":\"#173177\"},\"remark\":{\"value\":\"祝您出行愉快！\",\"color\":\"#173177\"}}}"
		payload := strings.NewReader(data)

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		response, e0 := client.Post(url, "application/json", payload)
		if e0 == nil {
			BodyBytes0, e1 := ioutil.ReadAll(response.Body)
			if e1 == nil {
				fmt.Println("发送消息值班人员结果", string(BodyBytes0))
				var Cal SendMsgCallBack
				Cal.ErrorCode = 100

				err11 := json.Unmarshal(BodyBytes0, &Cal)
				if err11 != nil {
					return
				}
				if Cal.ErrorCode == 0 {
					dao.Imp.OverMsg(&model.OverMsg{
						Id:     Interface2Int(msg["ID"]),
						MsgNum: 1,
					})
				}
			}
		}
	}

}
func SendMsg(msg map[string]interface{}) {
	templateid := "o2yh8P7T9K3rR9f4H_DxGpwYQntM4b3ZCv3EUnfwTQs"
	//url = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send?access_token=" + Token
	//data := "{\"touser\":\"" + msg["wechartid"].(string) + "\",\"mp_template_msg\":{\"appid\":\"wxa806018a131603d3\",\"template_id\":\"" + templateid + "\",\"url\":\"\",\"topcolor\":\"#FF0000\",\"miniprogram\":{\"appid\":\"wx032125bd60fe9474\",\"pagepath\":\"\"},\"data\":{\"first\":{\"value\":\"尊敬的 " + msg["CarNo"].(string) + " 车主，您已停车两个小时：\",\"color\":\"#173177\"},\"keyword1\":{\"value\":\"济南东高速服务区危化品车辆停车场\",\"color\":\"#173177\"},\"keyword2\":{\"value\":\"" + msg["intime"].(time.Time).String()[0:19] + "\",\"color\":\"#173177\"},\"keyword3\":{\"value\":\"--\",\"color\":\"#173177\"},\"keyword4\":{\"value\":\"--\",\"color\":\"#173177\"},\"keyword5\":{\"value\":\"--\",\"color\":\"#173177\"},\"remark\":{\"value\":\"祝您出行愉快！\",\"color\":\"#173177\"}}}}"

	//url := "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + Token
	url := "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=" + Token

	data := "{\"touser\":\"" + msg["wechartid"].(string) + "\",\"template_id\":\"" + templateid + "\",\"data\":{\"first\":{\"value\":\"尊敬的 " + msg["CarNo"].(string) + " 车主，您已停车超过两个小时，请尽快驶离！\",\"color\":\"#173177\"},\"keyword1\":{\"value\":\"济南东高速服务区危化品车辆停车场\",\"color\":\"#173177\"},\"keyword2\":{\"value\":\"" + msg["intime"].(time.Time).String()[0:19] + "\",\"color\":\"#173177\"},\"keyword3\":{\"value\":\"--\",\"color\":\"#173177\"},\"keyword4\":{\"value\":\"--\",\"color\":\"#173177\"},\"keyword5\":{\"value\":\"--\",\"color\":\"#173177\"},\"remark\":{\"value\":\"祝您出行愉快！\",\"color\":\"#173177\"}}}"
	payload := strings.NewReader(data)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, e0 := client.Post(url, "application/json", payload)
	if e0 == nil {
		BodyBytes0, e1 := ioutil.ReadAll(response.Body)
		if e1 == nil {
			fmt.Println("发送消息结果", string(BodyBytes0))
			var Cal SendMsgCallBack
			Cal.ErrorCode = 100

			err11 := json.Unmarshal(BodyBytes0, &Cal)
			if err11 != nil {
				return
			}
			if Cal.ErrorCode == 0 {
				dao.Imp.OverMsg(&model.OverMsg{
					Id:     Interface2Int(msg["ID"]),
					MsgNum: 1,
				})
			}
		}
	}
	//SendMsgToAdmin(msg, dao.Imp.GetAdminList())
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
