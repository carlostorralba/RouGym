package rutina

type Ejercicio struct {
	Nombre             string
	Series             int
	Repeticiones       int
	DescansoRepeticion int
}

func validacionEjercicio(ejercicio Ejercicio) []string {

	errores := []string{}

	if ejercicio.Nombre == "" {
		errores = append(errores, "Debe introducir un nombre de ejercicio")
	}
	if ejercicio.Series < 0 || ejercicio.Series > 10 {
		errores = append(errores, "Las series deben estar entre 1-10 series")
	}
	if ejercicio.Repeticiones < 0 || ejercicio.Repeticiones > 50 {
		errores = append(errores, "Las repeticiones deben estar entre 1-10 series")
	}
	if ejercicio.DescansoRepeticion < 30 || ejercicio.DescansoRepeticion > 300 {
		errores = append(errores, "El descanso por repetición debe estar entre 30-300 segundos")
	}

	return errores
}
