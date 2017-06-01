package server

import (
	"./api"

	"github.com/julienschmidt/httprouter"
	"net/http"
	"gopkg.in/unrolled/render.v1"
)

func SetupRouter() *httprouter.Router {
	r := httprouter.New()

	c := &api.BaseApiController{Render: render.New(render.Options{})}

	// Posts collection
	r.GET("/posts", c.PostsIndexHandler)
	r.POST("/posts", c.PostsCreateHandler)

	// Posts singular
	r.GET("/posts/:id", api.PostShowHandler)
	r.PUT("/posts/:id", api.PostUpdateHandler)
	r.GET("/posts/:id/edit", api.PostEditHandler)

	r.POST("/markdown", api.GenerateMarkdown)

	r.NotFound = http.FileServer(http.Dir("public"))

	return r
}

