package handlers

import (
	"chitchat/models"
	"fmt"
	"net/http"
)

func Login(write http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("auth.layout", "navbar", "login")
	t.Execute(write, nil)
}

func Signup(write http.ResponseWriter, request *http.Request) {
	generateHTML(write, nil, "auth.layout", "navbar", "signup")
}

func SignupAccount(write http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		fmt.Println("Cannot parse form")
	}

	user := models.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}

	if err := user.Create(); err != nil {
		fmt.Println(err)
		fmt.Println("Cannot create user")
	}

	http.Redirect(write, request, "/login", 302)
}

// POST /authenticate
// 通过邮箱和密码字段对用户进行认证
func Authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := models.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		fmt.Println("Cannot find user")
	}
	if user.Password == models.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			fmt.Println("Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}

// GET /logout
// 用户退出
func Logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		fmt.Println("Failed to get cookie")
		session := models.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", 302)
}
