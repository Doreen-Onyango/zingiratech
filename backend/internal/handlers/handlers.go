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
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "login.page.html", nil)
}
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "signup.page.html", nil)
}
func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "dashboard.page.html", nil)
}
func SchedulePickupHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "pickup.page.html", nil)
}
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	utils.RenderTemplate(w, "404.page.html", nil)
}

func UnauthorizedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	utils.RenderTemplate(w, "401.page.html", nil)
}

// ForbiddenHandler sends a 403 Forbidden response.
func ForbiddenHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	utils.RenderTemplate(w, "403.page.html", nil)
}

// InternalServerHandler sends a 500 Internal Server Error response.
func InternalServerHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	utils.RenderTemplate(w, "500.page.html", nil)
}
