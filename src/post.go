package blog

import (
	"strings"
	"time"
)

type Post struct {
	Subject  string
	Content  string
	Key      string
	Created  time.Time
	Last_Mod time.Time
}

func generate_key(p string) string {
	key := strings.Replace(p, " ", "_", -1)
	key = strings.ToLower(key)

	// TODO:
	// I really want to limit this to like 80 chars or somthing but I 
	// havent found anything yet to do  that easily

	return key
}

func NewPost(subject, content string) Post {
	return Post{Subject: subject,
		Content:  content,
		Key:      generate_key(subject),
		Created:  time.Now(),
		Last_Mod: time.Now()}
}
