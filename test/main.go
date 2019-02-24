package main

import (
	"io"
	"net/http"
)
import (
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()
	r.HandleFunc("/", RootHandler)
	http.Handle("/", r)
	http.ListenAndServe("0.0.0.0:8080", r)
}
func RootHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello World!")
	logrus.Info(request.RemoteAddr)
}
