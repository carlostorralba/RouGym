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

	/* 	ejercicio := rutina.Ejercicio{
	   		Nombre:             "Remo con barra",
	   		Series:             4,
	   		Repeticiones:       12,
	   		DescansoRepeticion: 90,
	   	}

	   	entrenamiento := rutina.Entrenamiento{
	   		Nombre:     "Espalda",
	   		Ejercicios: []rutina.Ejercicio{},
	   	}

	   	entrenamiento.Ejercicios = append(entrenamiento.Ejercicios, ejercicio)

	   	fmt.Println(entrenamiento.Ejercicios[0].Nombre)

	   	tipo := rutina.Ganar_Musculo

	   	rutina1 := rutina.Rutina{
	   		Nombre:             "Rutina 1",
	   		Dias_Entrenamiento: 4,
	   		Tipo:               tipo,
	   		Entrenamientos:     []rutina.Entrenamiento{},
	   		Duracion:           6,
	   	}

	   	rutina1.Entrenamientos = append(rutina1.Entrenamientos, entrenamiento)

	   	fmt.Println(rutina1.Entrenamientos[0].Nombre)
	   	fmt.Println(rutina1.Tipo)
	*/
	router := mux.NewRouter()

	router.HandleFunc("/rutinas", rutina.AddRutina).Methods("POST")

	fmt.Println("Escuchando por el puerto 5000")
	err := http.ListenAndServe(":5000", router)

	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
	fmt.Println("Escuchando por el puerto 5000")

}
