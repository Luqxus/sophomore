package api

import (
	"context"
	"log"
	"net/http"
	"time"
)

type apiFunc func(writer http.ResponseWriter, request *http.Request) error

type APIServer struct {
	listenAddress string
	router        *http.ServeMux
	handler       *handler
}

func New(listenAddress string) *APIServer {

	h := newHandler()

	return &APIServer{
		router:        http.NewServeMux(),
		listenAddress: listenAddress,
		handler:       h,
	}
}

func (api *APIServer) Run() error {

	// -- authentication --
	// user registration
	api.router.HandleFunc("POST /users/register", handlerFunc(api.handler.Register))

	// user login
	api.router.HandleFunc("POST /users/login", handlerFunc(api.handler.Login))

	// create new space
	api.router.HandleFunc("POST /spaces/create", handlerFunc(api.handler.CreateSpace))

	// run server and listen for incoming requests
	return http.ListenAndServe(api.listenAddress, api.router)
}

func handlerFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
		defer cancel()

		if err := fn(w, r.WithContext(ctx)); err != nil {
			log.Println(err.Error())
		}

	}
}
