package calculos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_soma(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     "positive numbers",
			a:        1,
			b:        1,
			expected: 2,
		},
		{
			name:     "negative numbers",
			a:        -1,
			b:        -2,
			expected: -3,
		},
		{
			name:     "mixed numbers",
			a:        10,
			b:        -1,
			expected: 9,
		},
		{
			name:     "tudo zero",
			a:        0,
			b:        0,
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Soma(tt.a, tt.b)
			assert.Equal(t, tt.expected, actual)
		})
	}

}
