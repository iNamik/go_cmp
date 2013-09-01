package cmp

// Uint8 is a cmp.T wrapper for basic type uint8
type Uint8 uint8

// Uint8::Cmp
func (a Uint8) Cmp(b_ interface{}) int {
	switch b := b_.(Uint8); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_Uint8 is a cmp.F for type cmp.Uint8
func F_Uint8(a_, b_ interface{}) int {
	switch a, b := a_.(Uint8), b_.(Uint8); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_uint8 is a cmp.F for basic type uint8
func F_uint8(a_, b_ interface{}) int {
	switch a, b := a_.(uint8), b_.(uint8); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// Compare_uint8 is a convenience function for
// explicity comparing basic type uint8
func Compare_uint8(a, b uint8) int {
	switch {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}
