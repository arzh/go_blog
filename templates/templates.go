package templates

import (
	"html/template"
)

var Pages map[string]*template.Template

// Pathing helper function
func p(filename string) string {
	return "templates/html/" + filename
}

func InitPages() error {
	Pages = make(map[string]*template.Template)
	var err error

	Pages["front"], err = template.ParseFiles(p("base.html"), p("front.html"), p("post.html"))
	if err != nil {
		return err
	}

	return err
}
