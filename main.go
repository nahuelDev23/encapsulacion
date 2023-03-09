
package main

import (
  "github.com/nahueldev23/encapsulacion/course"
)

func main(){
  Go := Course{
    Name:"Go desde cero",
    Price:12.34,
    IsFree:false,
    UserIDs:[]uint{12,13,14,15},
    Classes: map[uint]string{
      1:"Introduccion",
      2:"Estructuras",
      3:"Mapas",
    },
  }

 Go.PrintClasses()

}
