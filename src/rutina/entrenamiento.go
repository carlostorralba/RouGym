package rutina

type Entrenamiento struct {
	Nombre     string
	Ejercicios []Ejercicio
}

func validacionEntrenamiento(entrenamiento Entrenamiento) []string {
	errores := []string{}
	if entrenamiento.Nombre == "" {
		errores = append(errores, "Debe introducir un nombre de entrenamiento")
	}
	if len(entrenamiento.Ejercicios) < 0 || len(entrenamiento.Ejercicios) > 10 {
		errores = append(errores, "El número de ejercicios por entrenamiento tiene que ser entre 1-10")

	} else {
		for _, ejercicio := range entrenamiento.Ejercicios {
			errores = append(errores, validacionEjercicio(ejercicio)...)
		}
	}

	return errores
}
