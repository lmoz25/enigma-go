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

func (r *Rotor) Rotate() {
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
	if len(r.rotateCharacters) == 0 {
		r.Rotate()
	}
	return
}

type RotorSet struct {
	Rotors []*Rotor
}

func (rs *RotorSet) Encode(c rune) rune {
	// Init to something invalid
	lastRotorPosition := 'a'
	for _, rotor := range rs.Rotors {
		if _, ok := rotor.rotateCharacters[lastRotorPosition]; ok {
			rotor.Rotate()
		}
		c = rotor.encode(c)
		lastRotorPosition = rotor.position
	}
	return c
}
