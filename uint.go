package cmp

// Uint is a cmp.T wrapper for basic type uint
type Uint uint

// Uint::Cmp
func (a Uint) Cmp(b_ interface{}) int {
	switch b := b_.(Uint); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_Uint is a cmp.F for type cmp.Uint
func F_Uint(a_, b_ interface{}) int {
	switch a, b := a_.(Uint), b_.(Uint); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_uint is a cmp.F for basic type uint
func F_uint(a_, b_ interface{}) int {
	switch a, b := a_.(uint), b_.(uint); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// Compare_uint is a convenience function for
// explicity comparing basic type uint
func Compare_uint(a, b uint) int {
	switch {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}
