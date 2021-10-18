package main

import (
	"rougym/classes"
	"fmt"
)


func main()  {
	fmt.Printf("-------------------------------\n")
	fmt.Printf("-------------RouGym------------\n")
	fmt.Printf("-------------------------------\n")
	exercise := classes.Exercise {
		Name: "Remo con barra",
		Series: 4,
		Repetitions: 12,
		RestRepetition: 90,
	}

	training := classes.Training {
		Name: "Espalda",
		Exercises: []classes.Exercise{},
	}

	training.Exercises = append(training.Exercises, exercise)
	
	//fmt.Println(training.Exercises[0].Name)

	routine := classes.Routine {
		Name: "Rutina 1",
		TrainingDays: 4,
		Trainings: []classes.Training{},
	}

	routine.Trainings = append(routine.Trainings, training)

	//fmt.Println(routine.Trainings[0].Name)
}