package psql

import (
	"errors"
	"github.com/braddle/hex-arch-example-go/user"
	"math/rand"
)

type UserRepository struct {}

func (ur UserRepository) Save(u user.User) (string, error) {
	// TODO add randomness for response 1 happy, 2 DB Down
	i := rand.Intn(5)
	if i == 1 {
		return "", errors.New("DB is broken")
	} else {
		return "abc", nil
	}
}

func (ur UserRepository) GetUserById(id string) (user.User, error) {
	// Note: This would access the PSQL DB and populate the User for the response
	// TODO add randomness for response 1 happy, 2 no user, 3 DB Down
	return user.User{"Mark", "Bradley", 35, 5.11, true}, nil
}

func NewUserRepository() UserRepository {
	// Note: This would take a PSQL DB connection in the real world
	return UserRepository{}
}
