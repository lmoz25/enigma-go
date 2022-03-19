package components

type Encoder interface {
	Encode(rune) rune
}
