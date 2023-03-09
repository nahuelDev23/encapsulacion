package main

import (
	"fmt"

	"github.com/nahuelDev23/encapsulacion/course"
)

func main() {
	Go := course.NewCourse("GO DESDE CERO", 0, false)

	Go.UserIDs = []uint{12, 13, 14, 15}
	Go.SetClasses(map[uint]string{
		1: "Introduccion",
		2: "Estructuras",
		3: "Mapas",
	},
	)

	Go.SetName("POO CON GO")
	fmt.Println(Go.Name())

	Go.PrintClasses()
}

// Go.PrintClasses()
