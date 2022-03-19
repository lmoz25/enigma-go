package components

type Reflector struct {
	wiringTable map[rune]rune
}

func (r *Reflector) Encode(c rune) rune {
	return r.wiringTable[c]
}
