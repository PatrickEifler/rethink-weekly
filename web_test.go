package main

import (
	//"encoding/json"
	"bytes"
	r "github.com/dancannon/gorethink"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestSubscribe(t *testing.T) {
	setupTest()
	//defer teardownTest()

	res := httptest.NewRecorder()

	subscriber := map[string]string{
		"email":     "foo@foo.com",
		"firstname": "foo",
		"lastname":  "bar",
	}

	data := url.Values{}
	data.Set("email", subscriber["email"])
	data.Add("firstname", subscriber["firstname"])
	data.Add("lastname", subscriber["lastname"])

	req, _ := http.NewRequest("POST", "http://127.0.0.1:3000/api/subscriptions", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	h := SubscribeHandler()
	h.ServeHTTP(res, req)
	assert.Equal(t, res.Code, 200)

	result, _ := r.Table("subscribers").Filter(subscriber).Run(session)
	var existedSubscriber Subscriber
	result.One(&existedSubscriber)
	assert.Equal(t, existedSubscriber.Email, subscriber["email"], "Should get back the inserted supscription")
}

func T_estConfirmSubscription(t *testing.T) {
	setupTest()
	defer teardownTest()

	res := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "http://127.0.0.1:3000/api/subscriptions/1/confirm", nil)

	h := ConfirmSubscribeHandler()
	h.ServeHTTP(res, req)
	assert.Equal(t, res.Code, 200)

}
