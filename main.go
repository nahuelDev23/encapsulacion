
package main

import ("github.com/nahuelDev23/encapsulacion/course" 
"fmt")

func main(){
  Go := course.NewCourse("GO DESDE CERO",0,false)

  Go.UserIDs=[]uint{12,13,14,15}
  Go.Classes= map[uint]string{
      1:"Introduccion",
      2:"Estructuras",
      3:"Mapas",
    }

  fmt.Printf("%+v",Go)

  }

 // Go.PrintClasses()

