package cmp

// Uint16 is a cmp.T wrapper for basic type uint16
type Uint16 uint

// Uint16::Cmp
func (a Uint16) Cmp(b_ interface{}) int {
	switch b := b_.(Uint16); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_Uint16 is a cmp.F for type cmp.Uint16
func F_Uint16(a_, b_ interface{}) int {
	switch a, b := a_.(Uint16), b_.(Uint16); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_uint16 is a cmp.F for basic type uint16
func F_uint16(a_, b_ interface{}) int {
	switch a, b := a_.(uint16), b_.(uint16); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// Compare_uint16 is a convenience function for
// explicity comparing basic type uint16
func Compare_uint16(a, b uint16) int {
	switch {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}
