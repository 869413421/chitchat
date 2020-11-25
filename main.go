package main

import (
	. "chitchat/routes"
	"log"
	"net/http"
)

func main() {
	startWebServer("8989")
}

func startWebServer(port string) {
	r := NewRouter()
	http.Handle("/", r)

	// 处理静态资源文件
	assets := http.FileServer(http.Dir("public"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	log.Println("Starting Http Service at" + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
