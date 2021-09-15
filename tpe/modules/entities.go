package modules

import (
	"errors"
	"strconv"
	"unicode"
)

type Result struct {
	Type   string
	Value  string
	Length int
}

func NewResult(cadena string) (Result, error) {
	if cadenaValida(cadena) {
		tipo := cadena[0:2]
		largo, _ := strconv.ParseInt(cadena[2:4], 0, 8)
		valor := cadena[4:]
		return Result{tipo, valor, int(largo)}, nil
	} else {
		return Result{}, errors.New("cadena invalida")
	}

}

func cadenaValida(cadena string) bool {
	return len(cadena) > 4 && tipoValido(cadena) && largoValido(cadena)
}

func largoValido(cadena string) bool {
	largo, _ := strconv.ParseInt(cadena[2:4], 0, 8)
	valor := cadena[4:]
	return int(largo) == len(valor)
}

func tipoValido(cadena string) bool {
	tipo := cadena[0:2]
	valor := cadena[4:]
	if tipo == "TX" {
		return checkTx(valor)
	}
	if tipo == "NN" {
		return checkNumb(valor)
	}
	return false

}

func checkTx(valor string) bool {
	for _, c := range valor {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

func checkNumb(valor string) bool {
	for _, c := range valor {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
