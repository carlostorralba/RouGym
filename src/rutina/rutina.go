package rutina

type TipoRutina int

const (
	Definicion_Musculo TipoRutina = iota
	Ganar_Musculo
	Bajar_peso
	Resistencia_fisica

)

type Rutina struct {
	Nombre string
	Tipo TipoRutina
	Dias_Entrenamiento int
	Entrenamientos []Entrenamiento
	Duracion int
}
