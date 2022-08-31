package main

import (
	"fmt"
	"net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}
	//http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/static"))))
	http.HandleFunc("/", service.IndexHandler)

	http.HandleFunc("/api/count", service.CounterHandler)
	//log.Fatal(http.ListenAndServe(":80", http.FileServer(http.Dir("./static/"))))

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
	}

}
