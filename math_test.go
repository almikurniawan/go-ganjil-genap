package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckBilangan(t *testing.T) {
	assert := assert.New(t)
	bilangans := []struct {
		name     string
		bil      int
		expected string
	}{
		{
			name:     "uji01",
			bil:      4,
			expected: "Genap",
		},
		{
			name:     "uji02",
			bil:      7,
			expected: "Ganjil",
		},
	}

	for _, bilangan := range bilangans {
		t.Run(bilangan.name, func(t *testing.T) {
			assert.Equal(bilangan.expected, CheckBilangan(bilangan.bil))
		})
	}
}
