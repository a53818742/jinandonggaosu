package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func GetAllFile() {
	rd, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println("read dir fail:", err)
		return
	}
	fmt.Println("rd", rd)
	for _, fi := range rd {
		if !fi.IsDir() {
			fmt.Println(".......", fi.Name())
		}
	}

}

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	GetAllFile()
	//http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/static"))))
	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/MP_verify_X0kqrTo5XxsuQ4bB.txt", service.IndexText)
	http.HandleFunc("/index2.html", service.Index2Handler)
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
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
	}

}
