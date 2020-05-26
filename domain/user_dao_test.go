package domain

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	users = map[string]*User{
		"derekkenney@gmail.com": {
			ID:           123,
			FirstName:    "Derek",
			LastName:     "Kenney",
			Email:        "derekkenney@gmail.com",
			HashPassword: nil,
		},
	}
)

// For mocking the User functions for testing.
// Implements the Users interface
type mock struct{}

func (m *mock) Login(user *User) bool {
	if user == nil {
		return false
	}

	if user.ID != 123 {
		log.Printf("TEST: User id is incorrect. Can't login. Expected 123. Received %v\n", user.ID)
		return false
	}
	// Add mock for the Login call
	return true
}

// Looking at the user_dao domain object
// Create two test cases. One where we don't have any user matching
// a user id, and two where we do have a user matching a user id
func TestLoginIsTrue(t *testing.T) {

	// users mock
	m := new(mock)
	user := users["derekkenney@gmail.com"]
	loggedIn := m.Login(user)

	// Validation
	assert.NotNil(t, user, "We were expecting the user to not be nil")
	assert.EqualValues(t, "derekkenney@gmail.com", user.Email, "We were expecting the username to equal 'derekkenney@gmail.com'")
	assert.True(t, loggedIn, "Expected user to be logged in using email and password")
}

func TestLoginIsFalsel(t *testing.T) {

	// users mock
	m := new(mock)
	user := new(User)
	user.ID = 0
	loggedIn := m.Login(nil)

	// Validation
	assert.EqualValues(t, 0, 0, "We were expecting the user Id to be 0")
	assert.False(t, loggedIn, "Expected user login to be false")
}
