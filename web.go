package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	r "github.com/axcoto/rethink-weekly/Godeps/_workspace/src/github.com/dancannon/gorethink"
	"github.com/axcoto/rethink-weekly/Godeps/_workspace/src/github.com/getsentry/raven-go"
	"github.com/axcoto/rethink-weekly/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/axcoto/rethink-weekly/Godeps/_workspace/src/github.com/thoas/stats"
	"github.com/axcoto/rethink-weekly/Godeps/_workspace/src/github.com/unrolled/render"
)

type App struct // or "gopkg.in/unrolled/render.v1"
{
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

type Link struct {
	Id    string `gorethink:"id"`
	Title string `gorethink:"title"`
	Desc  string `gorethink:"desc"`
	Uri   string `gorethink:"uri"`
	Issue string `gorethink:"issue"`
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

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(out, "** middleware", r.URL)
		h.ServeHTTP(w, r)
	})
}

func run(mux *mux.Router) {
	listen := os.Getenv("LISTEN")
	fmt.Fprintf(out, "Server run on %s", listen)
	http.ListenAndServe(os.Getenv("LISTEN"), mux)
}

func runServer() {
	router := mux.NewRouter().StrictSlash(false)
	//r.HandleFunc("/", HomeHandler())
	router.HandleFunc("/api/issues", IssuesHandler())
	router.HandleFunc("/api/issues/{id:[0-9a-zA-Z-]+}", IssuesShowHandler())

	router.HandleFunc("/api/stats", StatsHandler())
	router.HandleFunc("/api/subscriptions", SubscribeHandler())
	router.HandleFunc("/api/subscriptions/{id:[0-9a-zA-Z-=_]+}", UnSubscribeHandler())
	router.HandleFunc("/api/subscriptions/{token:[0-9a-zA-Z-=_]+}/confirm", ConfirmSubscribeHandler())
	router.HandleFunc("/api/subscriptions/{token:[0-9a-zA-Z-=_]+}/ubsubscribe", UnSubscribeHandler())

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

	router.HandleFunc("/issues/{id:[0-9a-zA-Z-]+}", HomeHandler())
	router.HandleFunc("/about", HomeHandler())
	router.HandleFunc("/issues", HomeHandler())

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./client/dist/")))
	http.Handle("/", Middleware(router))

	run(router)
}

func HomeHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		//http.ServeFile(rw, req, r.URL.Path[1:])
		http.ServeFile(rw, req, "./client/dist/index.html")
	}
}

func StatsHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		res, err := r.Table("issues").Count().Run(session)
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
			fmt.Fprintf(out, "Err: %v\n", err)
		}
		if err = res.Err(); err != nil {
			raven.CaptureErrorAndWait(err, nil)
			fmt.Println(err)
		}

		var issueCount int
		err = res.One(&issueCount)
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
			fmt.Fprintf(out, "Err: %v\n", err)
		}
		res.Close()

		res, err = r.Table("subscribers").Count().Run(session)
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
			fmt.Fprintf(out, "Err: %v\n", err)
		}
		if err = res.Err(); err != nil {
			raven.CaptureErrorAndWait(err, nil)
			fmt.Println(err)
		}

		var subscriberCount int
		err = res.One(&subscriberCount)
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
			fmt.Fprintf(out, "Err: %v\n", err)
		}
		res.Close()

		body, err := json.Marshal(map[string]int{
			"issues":      issueCount,
			"subscribers": subscriberCount,
		})
		if err != nil {
			fmt.Fprintf(out, "Err: %v", err)
		}
		fmt.Fprintln(out, string(body))
		fmt.Fprintf(rw, string(body))
	}
}

func IssuesHandler() http.HandlerFunc {
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

func IssuesShowHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		res, err := r.Table("links").Filter(map[string]string{
			"issue": vars["id"],
		}).Run(session)
		fmt.Fprintf(out, "Get issue: %s", vars["id"])
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

func SubscribeHandler() http.HandlerFunc {
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
			res, err = r.Table("subscribers").Insert(map[string]string{
				"email":         subscriber.Email,
				"firstname":     subscriber.FirstName,
				"lastname":      subscriber.LastName,
				"status":        subscriber.Status,
				"confirm_token": subscriber.ConfirmToken,
			}).Run(session)
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
		if subscriber.Status == "pending" && subscriber.Email != "" {
			yeller.NotifiySubscriber(&subscriber)
		}
		fmt.Fprintf(rw, `{"result":"ok"}`)
	}
}

func ConfirmSubscribeHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		fmt.Fprintf(out, "db= %s", os.Getenv("RETHINK_DB"))

		var token string
		token = string(vars["token"])

		fmt.Fprintf(out, "\nConfirm token= %s\n", token)
		res, err := r.Table("subscribers").Filter(map[string]string{
			"confirm_token": token,
			"status":        "pending",
		}).Run(session)
		if err != nil {
			fmt.Fprintf(out, "Err: %v", err)
		}

		if err := res.Err(); err != nil {
			fmt.Println("Query error ", err)
		}

		var existedSubscriber Subscriber
		err = res.One(&existedSubscriber)
		if err != nil {
			fmt.Print("Error scanning database result: %s", err)
			fmt.Fprintf(rw, "%s", `{"result":"ok","message":"not found or approved"}`)
			return
		}

		res.Close()
		fmt.Println("Existed= %v", existedSubscriber)

		if existedSubscriber.Id != "" {
			r.Table("subscribers").Get(existedSubscriber.Id).Update(map[string]string{
				"status": "approved",
			}).Run(session)

			res, _ = r.Table("subscribers").Get(existedSubscriber.Id).Run(session)
			res.One(&existedSubscriber)

			yeller.ApproveSubscriber(&existedSubscriber)
			fmt.Fprintf(rw, "Cool, we all set. We will send you weekly update from now on :)")
		} else {
			fmt.Fprintf(rw, "%s", `{"result":"ok","message":"not found or approved"}`)
			fmt.Fprintf(out, "INvalid subscriber token")
		}
	}
}

func UnSubscribeHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		fmt.Fprintf(out, "db= %s", os.Getenv("RETHINK_DB"))

		var token string
		token = string(vars["token"])

		fmt.Fprintf(out, "\nUnsub token= %s\n", token)
		res, err := r.Table("subscribers").Filter(map[string]string{
			"confirm_token": token,
			"status":        "approved",
		}).Run(session)
		if err != nil {
			fmt.Fprintf(out, "Err: %v", err)
		}

		if err := res.Err(); err != nil {
			fmt.Println("Query error ", err)
		}

		var existedSubscriber Subscriber
		err = res.One(&existedSubscriber)
		if err != nil {
			fmt.Print("Error scanning database result: %s", err)
			fmt.Fprintf(rw, "%s", `{"result":"ok","message":"not found or approved"}`)
			return
		}

		res.Close()
		fmt.Println("Existed= %v", existedSubscriber)

		if existedSubscriber.Id != "" {
			r.Table("subscribers").Get(existedSubscriber.Id).Delete().Run(session)
			yeller.UnSubscribe(&existedSubscriber)
			fmt.Fprintf(rw, "We won't send you any email from now on")
		} else {
			fmt.Fprintf(rw, "%s", `{"result":"ok","message":"the subscriber isn't found"}`)
			fmt.Fprintf(out, "INvalid subscriber token")
		}
	}

}
