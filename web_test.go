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

func testDataSubscriber() map[string]string {
	return map[string]string{
		"email":     "foo@foo.com",
		"firstname": "foo",
		"lastname":  "bar",
	}
}

func testDataIssue() map[string]interface{} {
	var links [2]Link
	links[0] = Link{
		Uri:   "foo.com",
		Title: "foo",
		Desc:  "foo",
	}
	links[1] = Link{
		Uri:   "bar.com",
		Title: "bar",
		Desc:  "bar",
	}

	return map[string]interface{}{
		"name":  "Issue #1",
		"id":    "issue-1",
		"links": links,
	}
}

type MockedMail struct {
	mock.Mock
}

func (m *MockedMail) ApproveSubscriber(s *Subscriber) (bool, error) {
	m.Called(map[string]string{
		"status":    s.Status,
		"email":     s.Email,
		"firstname": s.FirstName,
		"lastname":  s.LastName,
	})

	return true, nil
}

func (m *MockedMail) UnSubscribe(s *Subscriber) (bool, error) {
	m.Called(map[string]string{
		"status":    s.Status,
		"email":     s.Email,
		"firstname": s.FirstName,
		"lastname":  s.LastName,
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

func TestIssue(t *testing.T) {
	setupTest()
	defer teardownTest()

	r.Table("issues").Insert(testDataIssue()).Run(session)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://127.0.0.1:3000/api/issues", nil)

	h := IssuesHandler()
	h.ServeHTTP(res, req)
	assert.Equal(t, res.Code, 200)
	//@TODO test JSON
	assert.Contains(t, res.Body.String(), "Issue #1")
	assert.Contains(t, res.Body.String(), "issue-1")
}

func TestShowIssue(t *testing.T) {
	setupTest()
	//defer teardownTest()

	r.Table("issues").Insert(testDataIssue()).Run(session)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://127.0.0.1:3000/api/issues/issue-1", nil)

	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/api/issues/{id:[0-9a-zA-Z-]+}", IssuesShowHandler())

	router.ServeHTTP(res, req)
	assert.Equal(t, res.Code, 200)
	//@TODO test JSON
	assert.Contains(t, res.Body.String(), "title")
	assert.Contains(t, res.Body.String(), "desc")
	assert.Contains(t, res.Body.String(), "foo.com")
	assert.Contains(t, res.Body.String(), "bar.com")
}

func TestSubscribe(t *testing.T) {
	setupTest()
	defer teardownTest()
	yeller = &MockedMail{}

	res := httptest.NewRecorder()

	data := url.Values{}
	data.Set("email", testDataSubscriber()["email"])
	data.Add("firstname", testDataSubscriber()["firstname"])
	data.Add("lastname", testDataSubscriber()["lastname"])

	req, _ := http.NewRequest("POST", "http://127.0.0.1:3000/api/subscriptions", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	yeller.(*MockedMail).On("NotifiySubscriber", testDataSubscriber()).Return(true, nil)

	h := SubscribeHandler()
	h.ServeHTTP(res, req)
	assert.Equal(t, res.Code, 200)

	result, _ := r.Table("subscribers").Filter(testDataSubscriber()).Run(session)
	var existedSubscriber Subscriber
	result.One(&existedSubscriber)
	assert.Equal(t, existedSubscriber.Email, testDataSubscriber()["email"], "Should get back the inserted supscription")
	yeller.(*MockedMail).AssertExpectations(t)
}

func TestConfirmSubscription(t *testing.T) {
	setupTest()
	defer teardownTest()
	yeller = &MockedMail{}

	newSubscriber := testDataSubscriber()
	newSubscriber["confirm_token"] = "i3khicon"
	newSubscriber["id"] = "i<3_khi_con"
	newSubscriber["status"] = "pending"

	r.Table("subscribers").Insert(newSubscriber).Run(session)

	res := httptest.NewRecorder()

	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/api/subscriptions/{token:[0-9a-zA-Z-=_]+}/confirm", ConfirmSubscribeHandler())

	approveSubscriber := testDataSubscriber()
	approveSubscriber["status"] = "approved"
	yeller.(*MockedMail).On("ApproveSubscriber", approveSubscriber).Return(true, nil)
	req, _ := http.NewRequest("GET", "http://127.0.0.1:3000/api/subscriptions/i3khicon/confirm", nil)

	router.ServeHTTP(res, req)
	assert.Equal(t, res.Code, 200)
	assert.Contains(t, res.Body.String(), "we all set")
	yeller.(*MockedMail).AssertExpectations(t)
}

func TestUbSubscribeHandler(t *testing.T) {
	setupTest()
	defer teardownTest()
	yeller = &MockedMail{}

	newSubscriber := testDataSubscriber()
	newSubscriber["confirm_token"] = "i3khicon"
	newSubscriber["id"] = "i<3_khi_con"
	newSubscriber["status"] = "approved"

	r.Table("subscribers").Insert(newSubscriber).Run(session)

	res := httptest.NewRecorder()

	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/api/subscriptions/{token:[0-9a-zA-Z-=_]+}/ubsubscribe", UnSubscribeHandler())

	approveSubscriber := testDataSubscriber()
	approveSubscriber["status"] = "approved"
	yeller.(*MockedMail).On("UnSubscribe", approveSubscriber).Return(true, nil)
	req, _ := http.NewRequest("GET", "http://127.0.0.1:3000/api/subscriptions/i3khicon/ubsubscribe", nil)

	router.ServeHTTP(res, req)
	assert.Equal(t, res.Code, 200)
	assert.Contains(t, res.Body.String(), "We won't send you any email from now on")
	yeller.(*MockedMail).AssertExpectations(t)
}
