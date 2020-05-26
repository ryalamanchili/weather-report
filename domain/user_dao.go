package domain

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/derekkenney/weather-report/common"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	err           error
	mongoDBClient *mongo.Client
)

// Users is an interface to allow the creation of a mock for use by
// the domain package
type Users interface {
	Login(user *User) bool
}

func (u *User) Login() bool {

	// Validate the user type properties
	if u.ID == 0 {
		log.Println("Invalid user id. Returned 0")
		return false
	}

	// implement the DA logic for login
	// create an instance of mongodb client
	mongoDBClient := common.GetMongoSession()

	// connect to collection
	collection := mongoDBClient.Database("weather").Collection("users")

	// create a filter to query for the user by email
	filter := bson.M{"email": u.Email}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = collection.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		log.Printf("Error validating user login for %s. %v\n", u.Email, err)
		return false
	}
	return true
}
