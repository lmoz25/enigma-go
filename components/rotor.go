package components

import (
	"fmt"

	"github.com/monzo/terrors"
)

type Rotor struct {
	position         rune
	wiringTable      map[rune]rune
	rotateCharacters map[rune]interface{}
}

// Use rotateCharacters = nil for a rotor that rotates after every key press
func NewRotor(startPosition rune, rotateCharacters []rune, rotorNumber int, wiringTable map[rune]rune) (*Rotor, error) {
	rotateCharsMap := map[rune]interface{}{}
	for _, c := range rotateCharacters {
		rotateCharsMap[c] = nil
	}
	wiring := wiringTable
	if wiring == nil {
		if rotorNumber > len(defaultRotors) {
			rotorNumberString := fmt.Sprintf("%d", rotorNumber+1)
			return nil, terrors.NotFound("default_wiring", "Default wiring map for rotor not found", map[string]string{
				"rotor_number": rotorNumberString,
			})
		}
		wiringTable = defaultRotors[rotorNumber]
	}

	return &Rotor{
		position:         startPosition,
		rotateCharacters: rotateCharsMap,
		wiringTable:      wiringTable,
	}, nil
}

func (r *Rotor) rotate() {
	r.position++
	if r.position > 'Z' {
		r.position = 'A'
	}
}

func (r *Rotor) encode(c rune) (encoded rune) {
	diff := r.position - 'A'
	fmt.Println(diff)
	leftHand := c + diff
	if leftHand > 'Z' {
		diff = diff - ('Z' - c + 1)
		leftHand = 'A' + diff
	}
	encoded = r.wiringTable[leftHand]
	// Use '*' to just check for rotors that rotate with every keypress
	if shouldRotate(r, '*') {
		r.rotate()
	}
	return
}

type RotorSet struct {
	Rotors     []*Rotor
	reflection bool
}

func (rs *RotorSet) Encode(c rune) rune {
	for i, rotor := range rs.Rotors {
		c = rotor.encode(c)
		if i+1 < len(rs.Rotors) && shouldRotate(rs.Rotors[i+1], rotor.position) {
			rs.Rotors[i+1].rotate()
		}
	}
	return c
}

func shouldRotate(rotor *Rotor, position rune) bool {
	_, ok := rotor.rotateCharacters[position]
	return ok || len(rotor.rotateCharacters) == 0
}
