package v1

import (
    "net/http"
)

func Authentication(h http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        h.ServeHTTP(w, r)
    }
    return http.HandlerFunc(fn)
}