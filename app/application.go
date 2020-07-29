package app

import (
	"github.com/braddle/hex-arch-example-go/psql"
	"github.com/braddle/hex-arch-example-go/queue"
	"github.com/gorilla/mux"
	"github.com/braddle/hex-arch-example-go/http/rest"
	"github.com/braddle/hex-arch-example-go/user"
	"net/http"
	"time"
)

func NewMicroservice() *Microservice {
	m := &Microservice{}
	m.initRouting()

	return m
}

type Microservice struct {
	r *mux.Router
}

func (m *Microservice) initRouting() {
	m.r = mux.NewRouter()
	m.r.Handle("/user/{id}", m.userGetHandle()).Methods(http.MethodGet)
	m.r.Handle("/user/", m.userPostHandle()).Methods(http.MethodPost)
}

func (m *Microservice) Serve(host string) error {
	srv := http.Server{
		Addr:         host,
		Handler:      m.r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return srv.ListenAndServe()
}

func (m *Microservice) userGetHandle() http.Handler {
	// TODO DI!
	return rest.NewUserGet(user.NewService(psql.NewUserRepository(), queue.NewKafka()))
}

func (m *Microservice) userPostHandle() http.Handler {
	// TODO DI!
	return rest.NewUserPost(user.NewService(psql.NewUserRepository(), queue.NewKafka()))
}







