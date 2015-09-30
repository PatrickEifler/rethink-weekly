package main

import (
	r "github.com/dancannon/gorethink"
	"log"
	"os"
)

func setupTest() {
	os.Setenv("RETHINK_DB", "test")

	var err error
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

	r.DB("test").TableCreate("issues").Run(session)
	r.DB("test").TableCreate("subscribers").Run(session)

	out = os.Stdout

	yeller = &mailer{}
}

func teardownTest() {

}
