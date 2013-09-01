package cmp

// Uint64 is a cmp.T wrapper for basic type uint64
type Uint64 uint64

// Uint64::Cmp
func (a Uint64) Cmp(b_ interface{}) int {
	switch b := b_.(Uint64); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_Uint64 is a cmp.F for type cmp.Uint64
func F_Uint64(a_, b_ interface{}) int {
	switch a, b := a_.(Uint64), b_.(Uint64); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_uint64 is a cmp.F for basic type uint64
func F_uint64(a_, b_ interface{}) int {
	switch a, b := a_.(uint64), b_.(uint64); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// Compare_uint64 is a convenience function for
// explicity comparing basic type uint64
func Compare_uint64(a, b uint64) int {
	switch {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}
