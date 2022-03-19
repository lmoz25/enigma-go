package components

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotor(t *testing.T) {
	startPos := 'D'
	rotor, err := NewRotor(startPos, nil, 0, nil)
	assert.NoError(t, err)

	assert.Equal(t, 'F', rotor.encode('A'))
	assert.Equal(t, startPos+1, rotor.position)

	assert.Equal(t, 'X', rotor.encode('M'))
	assert.Equal(t, startPos+2, rotor.position)

	// Test that rotor LHS encoding wraps back to A
	assert.Equal(t, 'F', rotor.encode('Y'))

	// Test that rotor position wraps round to A
	rotor.position = 'Z'
	rotor.encode('A')
	assert.Equal(t, 'A', rotor.position)
}

func TestRotorSet(t *testing.T) {
	startPos := 'O'
	rotor1, err := NewRotor(startPos, nil, 0, nil)
	assert.NoError(t, err)

	startPos = 'B'
	rotor2, err := NewRotor(startPos, []rune{'Q'}, 1, nil)
	assert.NoError(t, err)

	rotorSet := RotorSet{
		Rotors: []*Rotor{
			rotor1,
			rotor2,
		},
	}

	assert.Equal(t, 'E', rotorSet.Encode('A'))
	assert.Equal(t, startPos, rotor2.position)

	assert.Equal(t, 'B', rotorSet.Encode('A'))
	assert.Equal(t, startPos+1, rotor2.position)
}
