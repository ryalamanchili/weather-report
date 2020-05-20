package app

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/derekkenney/weather-report/controllers"
	"github.com/gorilla/mux"
)

// using asymetric RSA keys
// location of the files used for signing and verification
const (
	privateKeyPath = "./configs/app.rsa"
	publicKeyPath  = "./configs/app.rsa.pub"
)

var (
	verifyKey, signKey []byte
)

func StartApp() error {

	initKeys()

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
	router.HandleFunc("user/register", controllers.Register).Methods("POST")
	router.HandleFunc("user/login", controllers.Login).Methods("POST")
	return router
}

// Init the RSA keys before we start using the handlers
func initKeys() {
	var err error
	signKey, err = ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Printf("Error reading private key %v", err)
		return
	}

	verifyKey, err = ioutil.ReadFile(publicKeyPath)
	if err != nil {
		log.Printf("Error reading public key %v", err)
		return
	}
}
