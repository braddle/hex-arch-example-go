package rest

import (
	"encoding/json"
	"github.com/braddle/hex-arch-example-go/user"
	"io/ioutil"
	"net/http"
)

func NewUserPost(uc user.Creator) UserPost {
	return UserPost{uc}
}

type UserPost struct {
	uc user.Creator
}

func (h UserPost) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ju := JsonUser{}

	b, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(b, &ju)

	u := user.User{
		ju.Name,
		ju.Surname,
		ju.Age,
		ju.Height,
		ju.WearsGlasses,
	}

	id, err := h.uc.Create(u)

	if err != nil {
		if _, ok := err.(user.InvalidUser); ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		} else if _, ok := err.(user.StorageError); ok {
			w.WriteHeader(http.StatusBadGateway)
			w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		return
	}

	w.Header().Add("Location", "/user/" + id)
}
