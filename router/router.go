package router

import (
	"github.com/gorilla/mux"

	"net/http"

	PeerController "github.com/sjljrvis/g-bootstrap/controller"
)

// NewRouter is router pointer
func NewRouter() *mux.Router {
	fs := http.FileServer(http.Dir("./public/"))
	r := mux.NewRouter()

	/*
		user subrouter
		handle  REST-api /user here
	*/

	peerRouter := r.PathPrefix("/api/v1/peer").Subrouter()
	peerRouter.HandleFunc("/", PeerController.GetAll).Methods("GET")
	// userRouter.HandleFunc("/{id}", UserController.GetOne).Methods("GET")
	// userRouter.HandleFunc("/", UserController.Create).Methods("POST")

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
	return r
}
