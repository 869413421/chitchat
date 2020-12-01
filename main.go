package main

import (
	. "chitchat/config"
	. "chitchat/routes"
	"log"
	"net/http"
)

func main() {
	startWebServer("8989")
}

func startWebServer(port string) {
	// 在入口位置初始化全局配置
	config := LoadConfig()
	r := NewRouter()
	// 处理静态资源文件
	assets := http.FileServer(http.Dir(config.App.Static))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r)
	log.Println("Starting Http Service at " + port)
	err := http.ListenAndServe(config.App.Address, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
