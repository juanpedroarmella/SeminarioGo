package main

import (
	"fmt"
	"os"
)

//Funciones que empiecen en mayus->publicas, si empiezan en minusc->privadas
// Si son publicas se documentan con //**     */

type Person struct {
	Name string
	Age  int
}

type Logger func(string)

type IPerson interface {
	Describe(Logger) (string, error)
}

func NewPerson(n string, a int) Person {
	return Person{n, a}
}

func (p Person) Describe(l Logger) (string, error) {
	l("Logger:" + p.Name)
	return "Name: " + p.Name, nil
}

func main() {
	fmt.Printf(("hello, world\n"))
	var x int
	x = 10
	y := 20
	fmt.Println(x, y)

	arr := []int{}
	for i := 0; i < 4; i++ {
		arr = append(arr, i)
	}

	fmt.Println(arr)

	args := os.Args //argunmentos del main  go run main.go x y z por ej
	//range es como un foreach
	for _, k := range args {
		fmt.Println(k)
	}

	m := map[int]string{}
	m[0] = "cero"
	m[1] = "uno"
	m[1] = "dos"
	delete(m, 2)
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

	log := func(s string) {
		fmt.Println(s)
	}

	p := Person{"juan", 42}
	fmt.Println(p)

	fmt.Println(p.Describe(log))

	var ip IPerson
	ip = &Person{"pedro", 22}

	fmt.Println(ip.Describe(log))

	var p1 Person
	p1 = NewPerson("roberto", 40)

	fmt.Println(p1)

	n, err := ip.Describe(log)
	if err != nil {
		panic("se produjo un error")
	}

	fmt.Println(n)

}
