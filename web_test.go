package main

import (
	//"encoding/json"
	"bytes"
	r "github.com/dancannon/gorethink"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	subscriber map[string]string
)

func init() {
	subscriber = map[string]string{
		"email":     "foo@foo.com",
		"firstname": "foo",
		"lastname":  "bar",
	}
}

type MockedMail struct {
	mock.Mock
}

func (m *MockedMail) ApproveSubscriber(s *Subscriber) (bool, error) {
	m.Called(map[string]string{
		"id":            s.Id,
		"status":        s.Status,
		"confirm_token": s.ConfirmToken,
		"email":         s.Email,
		"firstname":     s.FirstName,
		"lastname":      s.LastName,
	})

	return true, nil
}

func (m *MockedMail) SendNewsletter() (int, error) {
	//m.Called()
	return 1, nil
}

func (m *MockedMail) NotifiySubscriber(s *Subscriber) (bool, error) {
	m.Called(map[string]string{
		"email":     s.Email,
		"firstname": s.FirstName,
		"lastname":  s.LastName,
	})
	return true, nil
}

func TestSubscribe(t *testing.T) {
	setupTest()
	defer teardownTest()
	yeller = &MockedMail{}

	res := httptest.NewRecorder()

	data := url.Values{}
	data.Set("email", subscriber["email"])
	data.Add("firstname", subscriber["firstname"])
	data.Add("lastname", subscriber["lastname"])

	req, _ := http.NewRequest("POST", "http://127.0.0.1:3000/api/subscriptions", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	yeller.(*MockedMail).On("NotifiySubscriber", subscriber).Return(true, nil)

	h := SubscribeHandler()
	h.ServeHTTP(res, req)
	assert.Equal(t, res.Code, 200)

	result, _ := r.Table("subscribers").Filter(subscriber).Run(session)
	var existedSubscriber Subscriber
	result.One(&existedSubscriber)
	assert.Equal(t, existedSubscriber.Email, subscriber["email"], "Should get back the inserted supscription")
	yeller.(*MockedMail).AssertExpectations(t)
}

func TestConfirmSubscription(t *testing.T) {
	setupTest()
	defer teardownTest()
	yeller = &MockedMail{}

	newSubscriber := subscriber
	newSubscriber["confirm_token"] = "i3khicon"
	newSubscriber["id"] = "i<3_khi_con"
	newSubscriber["status"] = "pending"

	r.Table("subscribers").Insert(newSubscriber).Run(session)

	res := httptest.NewRecorder()

	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/api/subscriptions/{token:[0-9a-zA-Z-=_]+}/confirm", ConfirmSubscribeHandler())

	yeller.(*MockedMail).On("ApproveSubscriber", subscriber).Return(true, nil)
	req, _ := http.NewRequest("GET", "http://127.0.0.1:3000/api/subscriptions/i3khicon/confirm", nil)

	router.ServeHTTP(res, req)
	assert.Equal(t, res.Code, 200)
	assert.Contains(t, res.Body.String(), "a")
	//yeller.(*MockedMail).AssertExpectations(t)
}
