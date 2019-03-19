package app

import (
	"github.com/gorilla/mux"
	"./APIs/"
)

func setUpRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/", APIs.)
}