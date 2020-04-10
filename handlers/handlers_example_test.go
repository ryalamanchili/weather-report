// Simple test to show how to write a basic Location request
package handlers_test

import (
	"encoding/json"
	"fmt"
	"github.com/derekkenney/doximity/handlers"
	"log"
	"net/http"
	"net/http/httptest"
)

func ExampleLocation() {
		router := handlers.Routes()
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/location/longitude/103/latitude/30", nil)
		router.ServeHTTP(rw, req)

		data := handlers.Data{}
		// Decode the JSON location from response

		if err := json.NewDecoder(rw.Body).Decode(&data); err != nil {
			log.Println("Error:", err)
		}
		fmt.Println(data)
		//Output:
		//{Kalamazoo 200}
}
