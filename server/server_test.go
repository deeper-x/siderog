package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafaeljusto/redigomock"
)

func TestStartSession(t *testing.T) {
	sess := Session{}
	mockConn := redigomock.NewConn()

	server := httptest.NewServer(sess.Start(mockConn))
	defer server.Close()

	expectedToken := "029384028095203892"
	urlQuery := fmt.Sprintf("%v/start_session?token=%v", server.URL, expectedToken)
	req, err := http.Get(urlQuery)

	if err != nil {
		log.Fatal(err)
	}

	result, err := ioutil.ReadAll(req.Body)
	req.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	output := string(result)

	if output != expectedToken {
		t.Errorf("output %v not expected", output)
	}
}
