package server

import (
	"assignment3/controllers"
	"log"
	"net/http"
)

func Start(PORT string) {
	http.HandleFunc("/", controllers.GetCuacaHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("views/assets/"))))
	log.Println("Server running at port", PORT)
	http.ListenAndServe(PORT, nil)
}
