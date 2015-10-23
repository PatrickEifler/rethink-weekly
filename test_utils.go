package main

import (
	r "github.com/axcoto/rethink-weekly/Godeps/_workspace/src/github.com/dancannon/gorethink"
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
	r.DB("test").TableCreate("links").Run(session)

	out = os.Stdout
}

func teardownTest() {
	r.DB("test").Table("issues").Delete().Run(session)
	r.DB("test").Table("subscribers").Delete().Run(session)
	r.DB("test").Table("links").Delete().Run(session)
}
