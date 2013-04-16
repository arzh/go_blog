package blog

import (
	//"fmt"
	"net/http"

	//"appengine"
	//"appengine/datastore"

	"pages"
)

func init() {
	pages.Init()
	http.HandleFunc("/", frontPage)
	//http.HandleFunc("/perma/([a-z+])", postPage)
}

func initData() []Post {

	post := NewPost("Hey new post", "This is a new post that I am making to test the DB")

	post2 := NewPost("Hey the second post", "This is a new post that I am making to test the DB AGAIN!")

	return []Post{post, post2}
}

func frontPage(w http.ResponseWriter, r *http.Request) {
	//context := appengine.NewContext(r)
	//InitPages()

	posts := initData()

	err := pages.Map["front"].Execute(w, posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//fmt.Fprintf(w, "Subject: %s | Content: %s | Created: %s", post_again.Subject, post_again.Content, post_again.Created)
}
