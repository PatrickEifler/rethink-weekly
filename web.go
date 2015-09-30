package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	r "github.com/dancannon/gorethink"
	"github.com/gorilla/mux"
	"github.com/unrolled/render" // or "gopkg.in/unrolled/render.v1"
	"net/http"
	//"os"
	"github.com/thoas/stats"
)

type App struct {
	R *render.Render
}

type Issue struct {
	Id   string `gorethink:"id"`
	Name string `gorethink:"name"`
}

type Subscriber struct {
	Id           string `gorethink:"id"`
	Email        string `gorethink:"email"`
	FirstName    string `gorethink:"firstname"`
	LastName     string `gorethink:"lastname"`
	Status       string `gorethink:"status"`
	ConfirmToken string `gorethink:"confirm_token"`
}

type Issues []*Issue

func template() *render.Render {
	r := render.New(render.Options{
		Layout:     "layout",
		Directory:  "templates",
		Charset:    "UTF-8",
		Extensions: []string{".tmpl", ".html"},
	})
	return r
}

type handlerAction func(rw http.ResponseWriter, req *http.Request)

func run(mux *mux.Router) {

	http.ListenAndServe("0.0.0.0:3000", mux)
}

func runServer() {
	router := mux.NewRouter().StrictSlash(false)
	//r.HandleFunc("/", HomeHandler())
	router.HandleFunc("/api/issues", IssuesHandler())
	router.HandleFunc("/api/issues/{id:[0-9a-zA-Z-]+}", IssuesShowHandler())

	router.HandleFunc("/api/subscriptions", SubscribeHandler())
	router.HandleFunc("/api/subscriptions/{id:[0-9a-zA-Z-=_]+}", UnSubscribeHandler())
	router.HandleFunc("/api/subscriptions/{token:[0-9a-zA-Z-=_]+}/confirm", ConfirmSubscribeHandler())
	router.HandleFunc("/api/subscriptions/{id:[0-9a-zA-Z-=_]+}/ubsubscribe", UnSubscribeHandler())

	stat := stats.New()
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		s, err := json.Marshal(stat.Data())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(s)
	})
	router.HandleFunc("/_stats", h)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./client/dist/")))
	fmt.Fprintf(out, "Server run on port 3000")
	run(router)
}

func HomeHandler() handlerAction {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Home")
	}
}

func IssuesHandler() handlerAction {
	return func(rw http.ResponseWriter, req *http.Request) {
		res, err := r.Table("issues").Run(session)
		if err != nil {
			fmt.Fprintf(out, "Err: %v\n", err)
		}

		if err = res.Err(); err != nil {
			fmt.Println(err)
		}

		defer res.Close()
		var rows []interface{}
		err = res.All(&rows)
		if err != nil {
			fmt.Fprintf(out, "Err: %v\n", err)
		}

		body, err := json.Marshal(rows)
		if err != nil {
			fmt.Fprintln(out, "Err: %v", err)
		}
		fmt.Fprintln(out, string(body))
		fmt.Fprintf(rw, string(body))
	}
}

func IssuesShowHandler() handlerAction {
	return func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		res, err := r.Table("issues").Get(vars["id"]).Run(session)
		if err != nil {
			fmt.Fprintf(out, "Err: %v\n", err)
		}

		if err = res.Err(); err != nil {
			fmt.Println(err)
		}

		defer res.Close()
		var row interface{}
		err = res.One(&row)
		if err != nil {
			fmt.Fprintf(out, "Err: %v\n", err)
		}

		body, err := json.Marshal(row)
		if err != nil {
			fmt.Fprintln(out, "Err: %v", err)
		}
		fmt.Fprintln(out, string(body))
		fmt.Fprintf(rw, string(body))
	}
}

func SubscribeHandler() handlerAction {
	return func(rw http.ResponseWriter, req *http.Request) {
		b := make([]byte, 32) // Get 32 random bytes
		_, err := rand.Read(b)

		if err != nil {
			fmt.Fprintln(out, "err %v", err)
		}

		token := base64.URLEncoding.EncodeToString(b)
		subscriber := Subscriber{
			Email:        req.FormValue("email"),
			FirstName:    req.FormValue("firstname"),
			LastName:     req.FormValue("lastname"),
			Status:       "pending",
			ConfirmToken: token,
		}

		res, err := r.Table("subscribers").Filter(map[string]interface{}{
			"email": subscriber.Email,
			//}).Count().Run(session)
		}).Run(session)
		if err != nil {
			fmt.Fprintf(out, "Err: %v", err)
		}
		var existedSubscriber Subscriber
		res.One(&existedSubscriber)
		res.Close()

		if existedSubscriber.Id != "" {
			fmt.Fprintf(out, "Existed email")
			subscriber = existedSubscriber
		} else {
			fmt.Fprintln(out, "subscribers: %v", subscriber)
			res, err = r.Table("subscribers").Insert(subscriber).Run(session)
			if err != nil {
				fmt.Fprintf(out, "Err: %v\n", err)
				fmt.Fprintf(rw, `{"result":"ok"}`)
				return
			}

			if err = res.Err(); err != nil {
				fmt.Fprintf(out, "Err Query: %v\n", err)
				fmt.Fprintf(rw, `{"result":"ok"}`)
				return
			}
		}

		//Notifi the subscribers
		if subscriber.Status == "pending" {
			yeller.NotifiySubscriber(&subscriber)
		}
		fmt.Fprintf(rw, `{"result":"ok"}`)
	}
}

func ConfirmSubscribeHandler() handlerAction {
	return func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		token := vars["token"]
		fmt.Println(out, "token= %s", token)
		res, err := r.Table("subscribers").Filter(map[string]string{
			"confirm_token": token,
		}).Run(session)
		if err != nil {
			fmt.Fprintf(out, "Err: %v", err)
		}
		var existedSubscriber Subscriber
		res.One(&existedSubscriber)
		res.Close()
		fmt.Println(existedSubscriber)

		if existedSubscriber.Id != "" {
			r.Table("subscribers").Get(existedSubscriber.Id).Update(map[string]string{
				"status": "approved",
			}).Run(session)

			yeller.ApproveSubscriber(&existedSubscriber)
		} else {
			fmt.Fprintf(out, "INvalid subscriber token")
		}
	}
}

func UnSubscribeHandler() handlerAction {
	return func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		fmt.Fprintln(rw, vars["id"])
	}
}
