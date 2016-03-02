package routes

import (
    "github.com/gorilla/mux"

    admin_controller "github.com/takaaki-mizuno/gorilla-boilerplate/app/http/controllers/admin"
)

func AdminRoutes(r *mux.Router) {

    s := r.PathPrefix("/admin").Subrouter()
    s.HandleFunc("/", admin_controller.IndexGetHandler).Methods("GET")

}
