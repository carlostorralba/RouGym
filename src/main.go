package main

import (
	"fmt"
	"net/http"
	"rougym/rutina"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Printf("-------------------------------\n")
	fmt.Printf("-------------RouGym------------\n")
	fmt.Printf("-------------------------------\n")

	router := mux.NewRouter()

	router.HandleFunc("/rutinas", rutina.AddRutina).Methods("POST")

	fmt.Println("Escuchando por el puerto 5000")
	err := http.ListenAndServe(":5000", router)

	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
	fmt.Println("Escuchando por el puerto 5000")

}
