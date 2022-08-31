package main

import (
	"fmt"
	"net/http"
)

func main() {
	//http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/static"))))

	err := http.ListenAndServe(":8080", http.FileServer(http.Dir("./static/")))
	if err != nil {
		fmt.Println(err)
	}

}
