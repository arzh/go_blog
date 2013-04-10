package blog

import (
	//"fmt"
	"html/template"
	"net/http"
	"time"

	//"appengine"
	//"appengine/datastore"
)

func init() {
	InitPages()
	http.HandleFunc("/", dbTester)
}

type PostDB struct {
	Subject  string
	Content  string
	Created  time.Time
	Last_Mod time.Time
}

func initData() []PostDB {

	post := PostDB{
		Subject:  "Hey new post",
		Content:  "This is a new post that I am making to test the DB",
		Created:  time.Now(),
		Last_Mod: time.Now()}

	post2 := PostDB{
		Subject:  "Hey the second post",
		Content:  "This is a new post that I am making to test the DB AGAIN!",
		Created:  time.Now(),
		Last_Mod: time.Now()}

	return []PostDB{post, post2}
}

var Pages map[string]*template.Template

func InitPages(w http.ResponseWriter) {
	p := "templates/"
	Pages = make(map[string]*template.Template)

	var err error
	Pages["front"], err = template.ParseFiles(p+"base.html", p+"front.html", p+"page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func dbTester(w http.ResponseWriter, r *http.Request) {
	//context := appengine.NewContext(r)
	InitPages()

	posts := initData()

	err := Pages["front"].Execute(w, posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//fmt.Fprintf(w, "Subject: %s | Content: %s | Created: %s", post_again.Subject, post_again.Content, post_again.Created)
}
