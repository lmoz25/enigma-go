package components

type Plugboard struct {
	wiringTable map[rune]rune
}

func (p *Plugboard) Encode(c rune) rune {
	encoded, found := p.wiringTable[c]
	if !found {
		return c
	}
	return encoded
}
