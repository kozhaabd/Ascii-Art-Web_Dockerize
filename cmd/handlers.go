package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"web/fs"
)

type Errors struct {
	Status  int
	Message string
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Errorhandler(w, 404)
		return
	}
	if r.Method != http.MethodGet {
		Errorhandler(w, 405)
		return
	}
	ts, err := template.ParseFiles("./html/Home.html")
	if err != nil {
		Errorhandler(w, 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		Errorhandler(w, 500)
	}
}

func Posthandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		Errorhandler(w, 404)
		return
	}
	if r.Method != http.MethodPost {
		Errorhandler(w, 405)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), 500)
		return
	}
	text := r.FormValue("text")

	banner := r.FormValue("font")
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		Errorhandler(w, 400)
		return
	}
	if text == "" || banner == "" {
		Errorhandler(w, 400)
		return
	}
	for _, v := range text {
		if (v < 32 || v > 126) && !(v == '\r' || v == '\n') {
			Errorhandler(w, 400)
			return
		}
	}

	var res string
	res = fs.Ascii(text, banner)
	html, err := template.ParseFiles("html/Home.html")
	if err != nil {
		Errorhandler(w, 500)
		return
	}
	err = html.Execute(w, res)
	if err != nil {
		Errorhandler(w, 500)
		return
	}
}

func Errorhandler(w http.ResponseWriter, status int) {
	html, err := template.ParseFiles("./html/Error.html")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "Internal Server Error 500")
		return
	}
	var Result Errors
	Result.Message = http.StatusText(status)
	Result.Status = status
	w.WriteHeader(status)
	err = html.Execute(w, Result)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "Internal Server Error 500")
		return
	}
}
