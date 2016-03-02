package routes

import (
    "github.com/gorilla/mux"

    user_controller "github.com/takaaki-mizuno/gorilla-boilerplate/app/http/controllers/user"
)

func UserRoutes(r *mux.Router) {

    s := r.PathPrefix("/").Subrouter()
    s.HandleFunc("/", user_controller.IndexGetHandler).Methods("GET")

}
