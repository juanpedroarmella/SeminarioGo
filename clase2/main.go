package main

import (
	"fmt"

	"tudai2021.com/model"
)

//le tenes que pasar un puntero sino hace una copia y no cambia el valor original
func changeName(p *model.Person, name string) {
	p.Name = name
}

func main() {
	p := model.NewPerson(0, "Daniela")
	changeName(&p, "Alicia")
	fmt.Println(p)
}
