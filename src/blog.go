package blog

import (
	//"fmt"
	"html/template"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
)

func init() {
	http.HandleFunc("/", dbTester)
}

type PostDB struct {
	Subject  string
	Content  string
	Created  time.Time
	Last_Mod time.Time
}

func dbTester(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)

	post := PostDB{
		Subject:  "Hey new post",
		Content:  "This is a new post that I am making to test the DB",
		Created:  time.Now(),
		Last_Mod: time.Now()}

	inc_key := datastore.NewIncompleteKey(context, "PostDB", nil)

	key, err := datastore.Put(context, inc_key, &post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var post_again PostDB
	if err := datastore.Get(context, key, &post_again); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	page, _ := template.ParseFiles("templates/blog.html")
	page.Execute(w, post_again)

	//fmt.Fprintf(w, "Subject: %s | Content: %s | Created: %s", post_again.Subject, post_again.Content, post_again.Created)
}
