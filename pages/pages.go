package pages

import (
	"html/template"
)

var Map map[string]*template.Template

// Pathing helper function
func p(filename string) string {
	return "pages/templates/" + filename
}

func Init() {
	Map = make(map[string]*template.Template)

	Map["front"] = template.Must(template.ParseFiles(p("base.html"), p("front.html"), p("post.html")))
	Map["newpost"] = template.Must(template.ParseFiles(p("base.html"), p("newpost.html")))
}
