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

	if len(validaciones) > 0 {
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

func validacionRutinas(rutina Rutina) []string {

	errores := []string{}

	if rutina.Nombre == "" {
		errores = append(errores, "Debe introducir un nombre de rutina")
	}
	if rutina.Dias_Entrenamiento < 0 || rutina.Dias_Entrenamiento > 7 {
		errores = append(errores, "Los días de entrenamiento tienen que ser entre 1-6 días")
	}
	if rutina.Tipo < 0 || rutina.Tipo > 4 {
		errores = append(errores, "El tipo de entrenamiento seleccionado no es válido")
	}
	if rutina.Duracion < 0 || rutina.Duracion > 13 {
		errores = append(errores, "La duración de la rutina tiene que estar entre 3-12 meses")
	}
	if len(rutina.Entrenamientos) != rutina.Dias_Entrenamiento {
		errores = append(errores, "El número de entrenamientos no correponde con los días de entrenamiento")
	} else {
		for _, entrenamiento := range rutina.Entrenamientos {
			errores = append(errores, validacionEntrenamiento(entrenamiento)...)
		}
	}

	return errores
}
