package main

import (
	"fmt"
	"net/http"

	"go-repository-pattern/controllers"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	response := controllers.ReadAllBanners
	fmt.Fprint(w, response)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/banners", handler).Methods("GET")
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8000", nil)
}
