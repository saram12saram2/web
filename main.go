package main

import (
	"fmt"
	"log"
	"net/http"
	"web/cmd"
)

// func main() {
// 	fmt.Println("Web server: http://localhost:8080/")
// 	http.HandleFunc("/", package.Home)
// 	http.ListenAndServe(":8080", nil)
// }

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", cmd.Home)
	mux.HandleFunc("/index.html", cmd.Create)
	log.Println("Запуск веб-сервера на http://localhost:8080/ ")
	fmt.Println("Server is listening...")
	fileServer := http.FileServer(http.Dir("./css/"))
	mux.Handle("/css/", http.StripPrefix("/css/", fileServer))
	err := http.ListenAndServe(":8085", mux)
	if err != nil {
		log.Fatal(err)
	}
}
