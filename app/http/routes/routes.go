package routes

import (
    "github.com/gorilla/mux"
    "github.com/justinas/alice"
    "github.com/takaaki-mizuno/gorilla-boilerplate/app/http/middlewares"
    "net/http"
)

func Routes() http.Handler {
    r := mux.NewRouter()
    StaticRoutes(r)

    APIV1Routes(r)
    AdminRoutes(r)
    UserRoutes(r)

    middleware := alice.New(middlewares.SecurityHeaders)

    return middleware.Then(r)
}

