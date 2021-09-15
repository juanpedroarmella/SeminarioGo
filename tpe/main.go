package main

import (
	"fmt"
	"tudai2021/modules"
)

func main() {
	cadena := "TX03ABC"
	result, err := modules.NewResult(cadena)
	if err == nil {
		fmt.Println(result)
	} else {
		panic(err)
	}

}
