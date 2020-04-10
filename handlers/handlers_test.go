// Test to show how to make an HTTP request to the location endpoint.
package handlers_test

import (
	"github.com/derekkenney/doximity/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

//TestLocation testing the location internal endpoint
func TestLocation(t *testing.T) {
	t.Log("Given the need to test the Location endpoint.")
	{
		router := handlers.Routes()
		rw := httptest.NewRecorder()

		req, err := http.NewRequest("GET", "/location/longitude/103/latitude/30", nil)
		if err != nil {
			t.Fatal("\tShould be able to create a request", ballotX, err)
		}
		t.Log("Should be able to create a request", checkMark)

		//Test that we get a status code 200 for a correct request
		router.ServeHTTP(rw, req)

		if rw.Code != 200 {
			t.Fatal("\tShould receive a \"200\"", ballotX, rw.Code)
		}
		t.Log("\tShould receive a \"200\"", checkMark)

		// Create a new request object with missing URL arguments
		}
}
func TestMissingCoordinates(t *testing.T) {
		t.Log("Given the need to test handling missing coordinates.")
		{
			router := handlers.Routes()
			rw := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/location/longitude/0/latitude/0", nil)
		router.ServeHTTP(rw, req)

		if rw.Code != 404 {
			t.Fatal("\tShould receive a \"404\" when missing coordinates", ballotX, rw.Code)
		}
		t.Log("Should receive a \"404\" when missing coordinates", checkMark)
	}
}
