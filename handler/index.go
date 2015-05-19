package handler

import (
	"github.com/unrolled/render" // or "gopkg.in/unrolled/render.v1"
	"net/http"
)

type Index struct {
	R *render.Render
}

func (c Index) Welcome(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Welcome, visit sub pages now."))
}

func (c Index) Data(w http.ResponseWriter, req *http.Request) {
	c.R.Data(w, http.StatusOK, []byte("Some binary data here."))
}

func (c Index) Json(w http.ResponseWriter, req *http.Request) {
	c.R.JSON(w, http.StatusOK, map[string]string{"hello": "json"})
}

func (c Index) Html(w http.ResponseWriter, req *http.Request) {
	// Assumes you have a template in ./templates called "example.tmpl"
	// $ mkdir -p templates && echo "<h1>Hello {{.}}.</h1>" > templates/example.tmpl
	c.R.HTML(w, http.StatusOK, "index/welcome", nil)
}
