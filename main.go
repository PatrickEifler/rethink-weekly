package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/unrolled/render" // or "gopkg.in/unrolled/render.v1"
	"net/http"
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

func run(mux *mux.Router) {
	http.ListenAndServe("0.0.0.0:3000", mux)
}

type handlerAction func(rw http.ResponseWriter, r *http.Request)

func main() {
	r := mux.NewRouter().StrictSlash(false)

	//r.HandleFunc("/", HomeHandler())
	r.HandleFunc("/api/issues/{id:[0-9]+}", IssuesHandler())
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./client/dist/")))
	run(r)
}

func HomeHandler() handlerAction {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Home")
	}
}

func IssuesHandler() handlerAction {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintln(rw, vars["id"])
	}
}

func SubscribeHandler() handlerAction {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintln(rw, vars["id"])
	}
}

func UnsubscribeHandler() handlerAction {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintln(rw, vars["id"])
	}
}
