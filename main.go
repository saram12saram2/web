package main

import (
	"fmt"
	"log"
	"net/http"
	"web/cmd"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", cmd.Home)
	mux.HandleFunc("/index.html", cmd.Create)
	log.Println("Запуск веб-сервера на http://localhost:8080/ ")
	fmt.Println("Server is listening...")
	fileServer := http.FileServer(http.Dir("./resources/"))
	mux.Handle("/resources/", http.StripPrefix("/resources/", fileServer))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
