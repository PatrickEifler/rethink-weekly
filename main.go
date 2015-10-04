package main

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	//"github.com/gorilla/mux"
	"log"
	//"github.com/rs/cors"
	//"github.com/thoas/stats"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var (
	session *r.Session
	out     io.Writer
	yeller  notifier
)

func main() {
	var err error
	yeller = &mailer{}

	session, err = r.Connect(r.ConnectOpts{
		Address:  os.Getenv("RETHINK_HOST"),
		Database: os.Getenv("RETHINK_DB"),
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

	var webCmd = &cobra.Command{
		Use:   "web",
		Short: "RethinkDB Weekly",
		Long:  `An application to send newsletter and allow subscribe to it `,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			runServer()
		},
	}

	var sendmailCmd = &cobra.Command{
		Use:   "send",
		Short: "Send out email for this issue",
		Long:  `Send out email for an issue. like: send issue-id-in-rethinkdb`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}

	var rootCmd = &cobra.Command{Use: "rewl"}
	rootCmd.AddCommand(webCmd, sendmailCmd)
	rootCmd.Execute()
}
