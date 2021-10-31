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

	almacenarRutinaJson()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("Rutina almacenada correctamente")

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
