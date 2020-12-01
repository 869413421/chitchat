package handlers

import (
	"chitchat/models"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
)

func NewThread(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	}
	generateHTML(w, nil, "layout", "auth.navbar", "new.thread")
}

// POST /thread/create
// 执行群组创建逻辑
func CreateThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			fmt.Println("Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			fmt.Println("Cannot find user")
			http.Redirect(writer, request, "/login", 302)
		}
		post := request.PostFormValue("post")
		_, err = user.CreateThread(post)
		if err != nil {
			fmt.Println("Cannot create thread")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// GET /thread/read
// 通过 ID 渲染指定群组页面
func ReadThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := models.ThreadByUUID(uuid)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "thread_not_found",
		})
		error_message(writer, request, msg)
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, &thread, "layout", "navbar", "thread")
		} else {
			generateHTML(writer, &thread, "layout", "auth.navbar", "auth.thread")
		}
	}
}
