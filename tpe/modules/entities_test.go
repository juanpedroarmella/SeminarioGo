package modules

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResult(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"NN04000A", false, "", "", 0},
	}

	for _, testData := range cases {
		r, err := NewResult(testData.Input)
		assert.Equal(t, err == nil, testData.Success)
		assert.Equal(t, r.Type, testData.Type)
		assert.Equal(t, r.Length, len(testData.Value))
		fmt.Println(err)

	}
}
