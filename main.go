package main

import (
	"log"
    "net/http"
    "github.com/gorilla/mux"

	"url-shortener/config"
	"url-shortener/handlers"
	"url-shortener/db"
)

func handle(conf config.HandlerTypeConf) {
	var handler func(http.ResponseWriter, *http.Request)

	if conf.Method == "GET" {
		handler = handlers.GetUrlHandler
	} else if conf.Method == "POST" {
		handler = handlers.ShortenUrlHandler
	} else{
		log.Fatal("ERROR, method not supported! Exiting..")
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(conf.Path, handler).Methods(conf.Method)
	log.Println("Starting server on port:", conf.Port)
	log.Fatal(http.ListenAndServe(":"+conf.Port, router))
}

func main() {
	var HandlerTypeConf config.HandlerTypeConf
	config.GetEnv(&HandlerTypeConf)

	err := db.InitConnections()
	if err != nil {
		log.Fatal("ERROR, could not connect to DB")
	}
	handle(HandlerTypeConf)
}
