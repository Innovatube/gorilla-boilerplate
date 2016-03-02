package v1

import (
    "net/http"

    "github.com/takaaki-mizuno/gorilla-boilerplate/app/services"
    api_v1_response "github.com/takaaki-mizuno/goji-boilerplate/app/http/responses/api/v1"

)

func StatusGetHandler(w http.ResponseWriter, r *http.Request) {
    services.Api().GenerateResponse(w, r, api_v1_response.Status{Status:"ok"})
}
