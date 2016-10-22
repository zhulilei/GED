package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/zhuapi/").Subrouter()
	s.HandleFunc("/v1/nginx/template", Nginxtemp).Methods("POST")
	http.Handle("/", r)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Println(err)
	}

}
