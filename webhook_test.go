package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetWebhookHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/webhook", nil)

	//if error is not nil
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("hub.verify_token", "token_verify")
	newRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(webhookGetHandler)

	handler.ServeHTTP(newRecorder, req)

	//check the status code if its what we're expecting

	if status := newRecorder.Code; status != http.StatusOK {
		t.Errorf("Our URL webhook GET URL handler returned the wrong status code. Got %v, want %v", status, http.StatusOK)
	}

	//convert our expected to string since in the
	//handler function, we wrote it as a slice of bytes
	expected := string([]byte("token_verify"))

	//check if our header contains the verify token

	if newRecorder.Body.String() != expected {
		t.Errorf("Hub token didn't match what is in the config file. Got: %v, expected: %s", newRecorder.Body.String(), expected)
	}
}
