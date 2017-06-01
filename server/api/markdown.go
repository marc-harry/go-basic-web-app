package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/russross/blackfriday"
)

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}