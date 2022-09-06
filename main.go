package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func GetAllFilePath(Path string) {
	rd, err := ioutil.ReadDir(Path)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return
	}
	for _, fi := range rd {
		if !fi.IsDir() {
			fmt.Println(".......", Path+fi.Name())
		} else {
			GetAllFilePath(Path + fi.Name() + "/")
		}
	}
}

func GetAllFile() {
	rd, err := ioutil.ReadDir("./")
	if err != nil {
		return
	}
	for _, fi := range rd {
		if !fi.IsDir() {
			fmt.Println("......./" + fi.Name())
		} else {
			GetAllFilePath("./" + fi.Name() + "/")
		}
	}
}

func GetToken() {

	url := "https://api.weixin.qq.com/cgi-bin/message/template/send"
	data := "{\"touser\":\"oCZvY55E1d6jDmo5wnGhvbZ4ROoo\",\"template_id\":\"o2yh8P7T9K3rR9f4H_DxGpwYQntM4b3ZCv3EUnfwTQs\",\"url\":\"\",\"topcolor\":\"#FF0000\",\"data\":{\"first\":{\"value\":\"尊敬的 京A00001 车主，您已停车两个小时：\",\"color\":\"#173177\"},\"keyword1\":{\"value\":\"济南东高速服务区危化品车辆停车场\",\"color\":\"#173177\"},\"keyword2\":{\"value\":\"2022-09-06 08:49:00\",\"color\":\"#173177\"},\"keyword3\":{\"value\":\"--\",\"color\":\"#173177\"},\"keyword4\":{\"value\":\"--\",\"color\":\"#173177\"},\"keyword5\":{\"value\":\"--\",\"color\":\"#173177\"},\"remark\":{\"value\":\"祝您出行愉快！\",\"color\":\"#173177\"}}}"
	payload := strings.NewReader(data)
	response, _ := http.Post(url, "application/json", payload)

	BodyBytes0, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(BodyBytes0))
}

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}
	GetToken()
	GetAllFile()
	//http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/static"))))
	http.HandleFunc("/", service.IndexHandler)
	//http.HandleFunc("/MP_verify_X0kqrTo5XxsuQ4bB.txt", service.IndexText)
	//http.HandleFunc("/index2.html", service.Index2Handler)
	http.HandleFunc("/api/CarUpdate", service.CarUpdate)
	http.HandleFunc("/api/CarInsert", service.CarInsert)
	http.HandleFunc("/api/CarOver", service.CarOver)
	http.HandleFunc("/api/CarGet", service.CarGet)
	http.HandleFunc("/api/CarList", service.CarList)
	http.HandleFunc("/api/GetUserInfo", service.GetUserInfo)

	http.HandleFunc("/api/AdminAdd", service.AdminAdd)
	http.HandleFunc("/api/AdminUpdate", service.AdminUpdate)
	http.HandleFunc("/api/AdminOver", service.AdminOver)
	http.HandleFunc("/api/AdminList", service.AdminList)

	http.HandleFunc("/api/CheckAdmin", service.CheckAdmin)
	http.HandleFunc("/api/UserLogin", service.UserLogin)

	//http.HandleFunc("/api/count", service.CounterHandler)
	//log.Fatal(http.ListenAndServe(":80", http.FileServer(http.Dir("./static/"))))

	//http.FileServer(http.Dir("./"))

	err := http.ListenAndServe(":80", http.FileServer(http.Dir("./")))
	if err != nil {
		fmt.Println(err)
	}

}
