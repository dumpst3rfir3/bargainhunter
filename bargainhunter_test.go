package bargainhunter_test

import (
	"bargainhunter"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFetchReturnsCorrectWebPageContent(t *testing.T) {
	var expectedtext = "Hello"
	testserver := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, expectedtext) }))
	bodytext, err := bargainhunter.Fetch(testserver.URL)
	if err != nil {
		t.Fatalf("Fetch returns an error: %s", err)
	}
	if !strings.Contains(bodytext, expectedtext) {
		t.Fatalf("Responexpectedtext,se body did not contain expected text %q. Body text: %q", expectedtext, bodytext)
	}
}

func TestInvalidUrlReturnsError(t *testing.T) {

	_, err := bargainhunter.Fetch("http://idonot.exist")
	if err == nil {
		t.Fatalf("Fetch should have returned an error, but the error was nil")
	}

}

func TestNotFoundUrlReturnsError(t *testing.T) {

	testserver := httptest.NewServer(nil)
	_, err := bargainhunter.Fetch(testserver.URL)
	if err == nil {
		t.Fatalf("Fetch should have returned an error, but the error was nil")
	}

}
