package routes

import (
    "github.com/gorilla/mux"
    "github.com/justinas/alice"

    api_v1_controllers "github.com/takaaki-mizuno/gorilla-boilerplate/app/http/controllers/api/v1"
    api_v1_middlewares "github.com/takaaki-mizuno/gorilla-boilerplate/app/http/middlewares/api/v1"
)

func APIV1Routes(r *mux.Router) {

    s := r.PathPrefix("/api/v1").Subrouter()
    middleware := alice.New(api_v1_middlewares.Authentication)

    s.Handle("/status", middleware.ThenFunc(api_v1_controllers.StatusGetHandler)).Methods("GET")

}
