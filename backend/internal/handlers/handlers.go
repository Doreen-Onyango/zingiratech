package handlers

import (
	"net/http"

	"github.com/Doreen-Onyango/zingiratech/backend/internal/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "Zingira Tech",
	}
	utils.RenderTemplate(w, "home.page.html", data)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "about.page.html", nil)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	utils.RenderTemplate(w, "404.page.html", nil)
}
