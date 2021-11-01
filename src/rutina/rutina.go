package rutina

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type TipoRutina int

const (
	Definicion_Musculo TipoRutina = iota
	Ganar_Musculo
	Bajar_peso
	Resistencia_fisica
)

type Rutina struct {
	Nombre             string
	Tipo               TipoRutina
	Dias_Entrenamiento int
	Entrenamientos     []Entrenamiento
	Duracion           int
}

var rutinas []Rutina = []Rutina{}

func AddRutina(w http.ResponseWriter, r *http.Request) {

	var rutina Rutina
	json.NewDecoder(r.Body).Decode(&rutina)

	rutinas = append(rutinas, rutina)

	w.Header().Set("Content-Type", "application/json")

	validaciones := validacionRutinas(rutina)

	if len(validaciones.Rutina) > 0 || len(validaciones.Entrenamientos) > 0 || len(validaciones.Ejercicios) > 0 {
		json.NewEncoder(w).Encode(validaciones)

	} else {
		almacenarRutinaJson()
		json.NewEncoder(w).Encode("Rutina almacenada correctamente")
	}

}

func almacenarRutinaJson() {
	jsonRutinas, err := json.Marshal(rutinas)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("rutinas.json", jsonRutinas, 0644)
	if err != nil {
		panic(err)
	}
}

func validacionRutinas(rutina Rutina) Error {

	var errores Error

	if rutina.Nombre == "" {
		errores.Rutina = append(errores.Rutina, "Debe introducir un nombre de rutina")
	}
	if rutina.Dias_Entrenamiento < 0 || rutina.Dias_Entrenamiento > 7 {
		errores.Rutina = append(errores.Rutina, "Los días de entrenamiento tienen que ser entre 1-6 días")
	}
	if rutina.Tipo < 0 || rutina.Tipo > 4 {
		errores.Rutina = append(errores.Rutina, "El tipo de entrenamiento seleccionado no es válido")
	}
	if rutina.Duracion < 0 || rutina.Duracion > 13 {
		errores.Rutina = append(errores.Rutina, "La duración de la rutina tiene que estar entre 3-12 meses")
	}
	if len(rutina.Entrenamientos) != rutina.Dias_Entrenamiento {
		errores.Rutina = append(errores.Rutina, "El número de entrenamientos no correponde con los días de entrenamiento")
	} else {
		for _, entrenamiento := range rutina.Entrenamientos {
			errores.Entrenamientos = append(errores.Entrenamientos, validacionEntrenamiento(entrenamiento)...)
			for _, ejercicio := range entrenamiento.Ejercicios {
				errores.Ejercicios = append(errores.Ejercicios, validacionEjercicio(ejercicio)...)
			}
		}
	}

	return errores
}
