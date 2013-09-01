package cmp

// Rune is a cmp.T wrapper for basic type rune
type Rune rune

// Rune::Cmp
func (a Rune) Cmp(b_ interface{}) int {
	switch b := b_.(Rune); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_Rune is a cmp.F for type cmp.Rune
func F_Rune(a_, b_ interface{}) int {
	switch a, b := a_.(Rune), b_.(Rune); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_rune is a cmp.F for basic type rune
func F_rune(a_, b_ interface{}) int {
	switch a, b := a_.(rune), b_.(rune); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// Compare_rune is a convenience function for
// explicity comparing basic type rune
func Compare_rune(a, b rune) int {
	switch {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}
