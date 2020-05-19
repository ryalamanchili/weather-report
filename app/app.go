package app

import (
	"log"
	"net/http"
	"time"

	"github.com/derekkenney/weather-report/controllers"
	"github.com/gorilla/mux"
)

// using asymetric RSA keys
// location of the files used for signing and verification
const (
	privateKeyPath = "../configs/app.rsa"
	publicKeyPath  = "../confgis/app.rsa.pub"
)

var (
	verifyKey, signKey []byte
)

func StartApp() error {

	log.Println("StartApp()")
	server := &http.Server{
		Addr:         ":8080",
		Handler:      initRoutes(),
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	log.Println("HTTP server is listening ...")
	return nil
}

func initRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/weather/location/{location/}", controllers.GetLocation).Methods("GET")
	router.HandleFunc("/user/authenticate", controllers.AuthenticateUser).Methods("POST")
	return router
}
