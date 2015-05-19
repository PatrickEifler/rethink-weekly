package main

import (
	"net/http"

	"github.com/unrolled/render" // or "gopkg.in/unrolled/render.v1"
	//"github.com/rs/cors"
	//"github.com/thoas/stats"
)

type App struct {
	R *render.Render
}

func template() *render.Render {
	r := render.New(render.Options{
		Layout:     "layout",
		Directory:  "templates",
		Charset:    "UTF-8",
		Extensions: []string{".tmpl", ".html"},
	})
	return r
}

func run(mux *http.ServeMux) {
	http.ListenAndServe("0.0.0.0:3000", mux)
}

func main() {
	mux := http.NewServeMux()
	r := template()
	app := App{
		r,
	}
	setupRoute(mux, &app)
	run(mux)
}
