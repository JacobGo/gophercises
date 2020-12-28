package urlshort

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMapHandler(t *testing.T) {


	rr := httptest.NewRecorder()
	mux := http.NewServeMux()

	pathToUrls := make(map[string]string)
	pathToUrls["/test"] = "http://www.example.com"

	handler := MapHandler(pathToUrls, mux)

	good, _ := http.NewRequest("GET", "/test", nil)
	handler.ServeHTTP(rr, good)
	if status := rr.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if location := rr.Header().Get("Location"); location != "http://www.example.com" {
		t.Errorf("handler returned wrong redirect lcoation: got %v want %v",
			location,"http://www.example.com" )
	}

	bad, _ := http.NewRequest("GET", "/bad", nil)
	handler.ServeHTTP(rr, bad)
	if status := rr.Code; status == http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}


}