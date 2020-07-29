package queue

import (
	"errors"
	"github.com/braddle/hex-arch-example-go/user"
	"math/rand"
)

func NewKafka() Kafka {
	return Kafka{}
}

type Kafka struct {}

func (k Kafka) UserCreated(id string, u user.User) error {
	if rand.Intn(2) == 1 {
		return errors.New("Kafka broken")
	}

	return nil
}
