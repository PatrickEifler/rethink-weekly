package main

import (
	"github.com/kureikain/rethink-weekly/handler"
	"net/http"
)

func setupRoute(mux *http.ServeMux, app *App) {
	r := template()

	mux.Handle("/static", http.FileServer(http.Dir("/")))

	indexHandler := handler.Index{r}
	mux.HandleFunc("/", indexHandler.Welcome)
	mux.HandleFunc("/data", indexHandler.Data)
	mux.HandleFunc("/json", indexHandler.Json)
	mux.HandleFunc("/html", indexHandler.Html)

}
