package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConfirmSubscription(t *testing.T) {
	setupTest()
	defer teardownTest()

	res := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "http://127.0.0.1:3000/api/subscriptions/1/confirm", nil)

	h := ConfirmSubscribeHandler()
	h.ServeHTTP(res, req)
	assert.Equal(t, res.Code, 200)

}
