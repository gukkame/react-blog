package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)

	database()

	//IMAGES -> ./resources
	fileServer := http.FileServer(http.Dir("./resources"))
	http.Handle("/resources/", http.StripPrefix("/resources", fileServer))
	//GOLANG SERVER
	fmt.Printf("API Server running at port http://localhost:8080/\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("(server.go) Golang server has stopped due to:")
		log.Fatal(err)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("JUP ERROR")
		panic(err)
	}
}
func home(http.ResponseWriter, *http.Request) {}
