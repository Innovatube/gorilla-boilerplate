package routes

import (
    "github.com/gorilla/mux"

    "net/http"
)

func StaticRoutes(r *mux.Router) {

    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

}
