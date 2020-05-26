package common

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// using asymetric RSA keys
// location of the files used for signing and verification
const (
	privateKeyPath = "./common/app.rsa"
	publicKeyPath  = "./common/app.rsa.pub"
I c)

var (
	AppConfig          appConfig
	verifyKey, signKey []byte
	mongoDBClient      *mongo.Client
	err                error
)

// Holds configuration values for the Weather service
// persistence layer
type dbConfig struct {
	MongoDBHost string
	DBUser      string
	DBPwd       string
}

type rpcConfig struct {
}

type appConfig struct {
	DBConfig  dbConfig
	RPCConfig rpcConfig
}

// Initialize AppConfig
func initConfig() {
	AppConfig.DBConfig.MongoDBHost = os.Getenv("MONGODB_HOST")
	AppConfig.DBConfig.DBUser = os.Getenv("DB_USER")
	AppConfig.DBConfig.DBPwd = os.Getenv("DB_PASSWORD")
}

// Init the RSA keys before we start using the
// We use the private key for signing JWTs, and public
// key verifies the JWT in an HTTP request
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

// GetMongoSession()
//
// Returns a copy of a MongoDB client. The client is
// configured using the DBConfig properties
func GetMongoSession() *mongo.Client {
	if mongoDBClient == nil {
		mongoDBClient, err = mongo.NewClient(options.Client().ApplyURI(AppConfig.DBConfig.MongoDBHost))

		if err != nil {
			log.Fatalf("Error creating a database client %s\n", err.Error())
		}

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		mongoDBClient.Connect(ctx)
	}

	return mongoDBClient
}
