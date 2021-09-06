package main

import (
	"testing"

	"tudai2021.com/model"

	"github.com/stretchr/testify/assert"
)

func TestChangeName(t *testing.T) {

	p := model.NewPerson(1, "Alicia")
	changeName(&p, "Agustin")

	assert.Equal(t, p.Name, "Agustin", "los valores no coinciden")
}
