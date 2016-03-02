package services

import (
    "github.com/ugorji/go/codec"
    "net/http"
)

type api struct {
}

func Api() *api {
    return &api{}
}

func (api *api) GenerateResponse(w http.ResponseWriter, r *http.Request, data interface{}) {

    if r.Header.Get("Accept") == "application/x-msgpack" {
        w.Header().Set("Content-Type", "application/x-msgpack")
        h := new(codec.MsgpackHandle)
        encoder := codec.NewEncoder(w, h)
        encoder.Encode(data)
    } else {
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Content-Security-Policy", "default-src 'none'")
        h := new(codec.JsonHandle)
        encoder := codec.NewEncoder(w, h)
        encoder.Encode(data)
    }

}