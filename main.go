package main

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	"github.com/gorilla/mux"
	"github.com/unrolled/render" // or "gopkg.in/unrolled/render.v1"
	"log"
	"net/http"
	//"github.com/rs/cors"
	//"github.com/thoas/stats"
	"encoding/json"
	"io"
	//"io/ioutil"
	"github.com/mailgun/mailgun-go"
	"github.com/spf13/cobra"
	"os"
)

type App struct {
	R *render.Render
}

type Issue struct {
	Id   string `gorethink:"id"`
	Name string `gorethink:"name"`
}

type Subscriber struct {
	Email     string `gorethink:"email"`
	FirstName string `gorethink:"firstname"`
	LastName  string `gorethink:"lastname"`
}

type Issues []*Issue

var (
	session *r.Session
	out     io.Writer
)

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

type handlerAction func(rw http.ResponseWriter, req *http.Request)

func main() {
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address:  "127.0.0.1:28015",
		Database: "rewl",
		MaxIdle:  20,
		MaxOpen:  20,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err != nil {
		log.Fatalln(err.Error())
	}
	session.SetMaxOpenConns(10)

	//out := ioutil.Discard
	out = os.Stdout

	var HugoCmd = &cobra.Command{
		Use:   "web",
		Short: "RethinkDB Weeklet",
		Long:  `An application to send newsletter and allow subscribe to it `,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			runServer()
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdPrint, cmdEcho)
	cmdEcho.AddCommand(cmdTimes)
	rootCmd.Execute()
}

func runServer() {
	router := mux.NewRouter().StrictSlash(false)
	//r.HandleFunc("/", HomeHandler())
	router.HandleFunc("/api/issues", IssuesHandler())
	router.HandleFunc("/api/issues/{id:[0-9a-zA-Z-]+}", IssuesShowHandler())
	router.HandleFunc("/api/subscriptions", SubscribeHandler())
	router.HandleFunc("/api/subscriptions/{id:[0-9a-zA-Z-]+}", UnSubscribeHandler())
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
		//vars := mux.Vars(req)
		subscriber := Subscriber{req.FormValue("email"), req.FormValue("firstname"), req.FormValue("lastname")}
		fmt.Fprintln(out, "subscribers: %v", subscriber)
		res, err := r.Table("subscribers").Insert(subscriber).Run(session)
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

		//Notifi the subscribers
		notifiySubscriber(&subscriber)
		fmt.Fprintf(rw, `{"result":"ok"}`)
	}
}

func UnSubscribeHandler() handlerAction {
	return func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		fmt.Fprintln(rw, vars["id"])
	}
}

func notifiySubscriber(subscriber *Subscriber) {
	fmt.Fprintln(out, "WIll send email to %s", subscriber.Email)
	mg := mailgun.NewMailgun("mg.noty.im", os.Getenv("MAILGUN_API"), "")

	m := mg.NewMessage(
		fmt.Sprintf("Vinh <%s>", os.Getenv("MAIL_FROM")),                                                                                                            // From
		"Verify subscriptions on RethinkDB Goodies",                                                                                                                 // Subject
		"Hi please follow this link <a href=\"http://127.0.0.1:3000/api/subscriont/1/verify\"> to verify</a>. <br />You can ignore if you don't want to subscribe.", // Plain-text body
		"kurei <kurei@axcoto.com>", // Recipients (vararg list)
	)

	_, _, err := mg.Send(m)

	if err != nil {
		log.Fatal(err)
	}
}

// Send out news letter on due date
func SendNewsletter() {

}
