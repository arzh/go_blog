package blog

import (
	//"fmt"
	"net/http"

	"appengine"
	//"appengine/datastore"
	"DB/Posts"
	"pages"
	"strings"
)

func init() {
	pages.Init()
	http.HandleFunc("/", frontPage)
	http.HandleFunc("/new", NewpostHandler)
	http.HandleFunc("/init", initData)
	http.HandleFunc("/p/", testPerma)
	//http.HandleFunc("/perma/([a-z+])", postPage)
}


func testPerma(w http.ResponseWriter, r *http.Request) {
	pathSplit := strings.Split(r.URL.Path, "/")
	key := pathSplit[len(pathSplit)-1:]
	//fmt.Fprint(w, pathSplit)
	//fmt.Fprint(w, key)
	
	p, err := posts.Get(appengine.NewContext(r), key[0])

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = pages.Map["permalink"].Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func initData(w http.ResponseWriter, r *http.Request) {
	post := posts.New("Hey new post", "This is a new post that I am making to test the DB")
	post2 := posts.New("Hey the second post", "This is a new post that I am making to test the DB AGAIN!")
	c := appengine.NewContext(r)
	post.Put(c)
	post2.Put(c)
}

func frontPage(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	lastesPosts, err := posts.GetLatest(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = pages.Map["front"].Execute(w, lastesPosts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
