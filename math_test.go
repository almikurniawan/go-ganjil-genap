package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckBilangan(t *testing.T) {
	assert := assert.New(t)
	bilangans := []int{4, 7, 8, 19, 189}
	expected := []string{"Genap", "Ganjil", "Genap", "Ganjil", "Ganjil"}

	results := CheckBilangan(bilangans...)

	for i := 0; i < len(results); i++ {
		assert.Equal(expected[i], results[i])
	}
}

func BenchmarkCheckBilangan(b *testing.B) {
	bilangans := []int{4, 7, 8, 19, 189}
	CheckBilangan(bilangans...)
}
