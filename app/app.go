package app

import (
    "github.com/takaaki-mizuno/gorilla-boilerplate/app/http/routes"
    "net/http"
)

func App() {
    r := routes.Routes()
    http.ListenAndServe(":8000", r)
}
