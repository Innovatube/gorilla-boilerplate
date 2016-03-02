package middlewares

import (
    "net/http"
)

func SecurityHeaders(h http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-XSS-Protection", "1; mode=block")
        w.Header().Set("X-UA-Compatible", "chrome=1")
        h.ServeHTTP(w, r)
    }
    return http.HandlerFunc(fn)
}