package rest

import (
	"encoding/json"
	"github.com/braddle/hex-arch-example-go/user"
	"net/http"
)

func NewUserGet(ur user.IDRetriever) UserGet {
	return UserGet{ur}
}

type UserGet struct {
	ur user.IDRetriever
}

type JsonUser struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	Age int `json:"age"`
	Height float32 `json:"height"`
	WearsGlasses bool `json:"wears_glasses"`
}

func (h UserGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, _ := h.ur.RetrieveUser("q")

	ju := JsonUser{
		u.FirstName,
		u.LastName,
		u.Age,
		u.Height,
		u.Glasses,
	}

	body, _ := json.Marshal(ju)
	w.Write(body)
}
