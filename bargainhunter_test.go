package bargainhunter_test

import (
	"bargainhunter"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestFetchReturnsCorrectWebPageContent(t *testing.T) {
	t.Parallel()
	var expectedtext = "Hello"
	testserver := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expectedtext)
	}))
	bodytext, err := bargainhunter.Fetch(testserver.URL)
	if err != nil {
		t.Fatalf("Fetch returns an error: %s", err)
	}
	if !strings.Contains(bodytext, expectedtext) {
		t.Fatalf("Response body did not contain expected text %q. Body text: %q", expectedtext, bodytext)
	}
}

func TestInvalidUrlReturnsError(t *testing.T) {
	t.Parallel()
	_, err := bargainhunter.Fetch("http://idonot.exist")
	if err == nil {
		t.Fatalf("Fetch should have returned an error, but the error was nil")
	}

}

func TestNotFoundUrlReturnsError(t *testing.T) {
	t.Parallel()
	testserver := httptest.NewServer(nil)
	_, err := bargainhunter.Fetch(testserver.URL)
	if err == nil {
		t.Fatalf("Fetch should have returned an error, but the error was nil")
	}
}

func TestExtractPrice_ReturnsPriceFromNeweggProductPageHTML(t *testing.T) {
	input, err := os.ReadFile("testdata/newegg.html")
	if err != nil {
		t.Fatal(err)
	}
	want := 499.99
	got, err := bargainhunter.ExtractPrice(string(input))
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("want %.2f, got %.2f", want, got)
	}
}

func TestExtractPrice_ReturnsErrorIfPriceNotFound(t *testing.T) {
	t.Parallel()
	_, err := bargainhunter.ExtractPrice("")
	if err == nil {
		t.Errorf("want error for missing price, got nil")
	}
}

func TestExtractPrice_ReturnsFloatParseError(t *testing.T) {
	t.Parallel()
	_, err := bargainhunter.ExtractPrice(`"FinalPrice":50.0.0,"Instock"`)
	if err == nil {
		t.Errorf("want error for missing price, got nil")
	}
}
