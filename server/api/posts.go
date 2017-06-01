package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

type Post struct {
	Id int `json:id`
	Title string `json:title`
	Message string `json:message`
}

func (c *BaseApiController) PostsIndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var posts [10]Post
	for i := 0; i < 10; i++ {
		posts[i] = Post{i,"Test"+strconv.Itoa(i), "Message + "+strconv.Itoa(i)}
	}

	c.JSON(rw, 200, posts)
}

func (c *BaseApiController) PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	posts := []Post{}

	json.NewDecoder(r.Body).Decode(&posts)

	for i := 0; i < len(posts); i++ {
		posts[i].Id = i
	}
	c.JSON(rw, 200, posts)
}

func PostShowHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Fprintln(rw, "showing post", id)
}

func PostUpdateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post update")
}

func PostDeleteHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post delete")
}

func PostEditHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "post edit")
}