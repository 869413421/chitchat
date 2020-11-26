package handlers

import (
	"chitchat/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	threads, err := models.Threads();
	if err == nil {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout", "navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "auth.navbar", "index")
		}
	}
}
