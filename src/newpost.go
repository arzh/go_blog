package blog

import (
	"pages"
	"DB/Posts"
	"appengine"
	"net/http"
)

type PageDef struct {
	Title string
	Text string

	Titleclass string
	Textclass string

	Titleerror string
	Texterror string
}

func emptyPage() PageDef {
	return PageDef{"", "", "", "", "", ""}
}

func render_page(w http.ResponseWriter, def PageDef){

	err := pages.Map["newpost"].Execute(w, def)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func NewpostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		render_page(w, emptyPage())
	} else if r.Method == "POST" {
		pageState := emptyPage()
		hasError := false
		pageState.Title = r.FormValue("subject")
		if len(pageState.Title) <= 0 {
			pageState.Titleclass = "error-box"
			pageState.Titleerror = "You need a title!"
			hasError = true
		}
		pageState.Text = r.FormValue("content")
		if len(pageState.Text) <= 0 {
			pageState.Textclass = "error-box"
			pageState.Texterror = "You need some content!"
			hasError = true
		}

		if hasError {
			render_page(w, pageState)
		} else {
			p := posts.New(pageState.Title, pageState.Text)
			p.Put(appengine.NewContext(r))
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}

	}
}
