package main

import (
	"fmt"
	"io/ioutil"
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

	//set our dummy query string the same way facebook sends
	//when verifying our verify_token
	q := req.URL.Query()
	q.Set("hub.verify_token", "token_verify")
	q.Set("hub.challenge", "abc123")
	req.URL.RawQuery = q.Encode()

	//log our encoded URL to troubleshoot failing test
	newRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(webhookGetHandler)

	handler.ServeHTTP(newRecorder, req)

	//check the status code if its what we're expecting

	if status := newRecorder.Code; status != http.StatusOK {
		t.Errorf("Our URL webhook GET URL handler returned the wrong status code. Got %v, want %v", status, http.StatusOK)
	}

	//convert our expected to string since in the
	//handler function, we wrote it as a slice of bytes
	expected := []byte("abc123")

	//check if our header contains the verify token
	bd, _ := ioutil.ReadAll(newRecorder.Body)
	fmt.Println(string(bd))
	if newRecorder.Body.Bytes() == nil {
		t.Errorf("Hub token didn't match what is in the config file. Got: %v, expected: %s", newRecorder.Body.String(), expected)
	}
}
