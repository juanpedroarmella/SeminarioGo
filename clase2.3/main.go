package main

import (
	"fmt"
	"strings"
)

func main() {
	country := "Japan"
	for i, c := range country {
		//d:pos c:car
		fmt.Printf("%d:%q", i, c)
	}
	fmt.Println()
	fmt.Println(country[1:3])
	fmt.Println(country[1:])

	fmt.Println(strings.Contains("Japan", "a"))

	fmt.Println(strings.Count("Japan", "a"))

	fmt.Println("Japan" < "a")
}
