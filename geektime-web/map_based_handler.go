package main

import (
	"net/http"
)

type Handler struct {
}

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}
