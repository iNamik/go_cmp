package cmp

// Uint32 is a cmp.T wrapper for basic type uint32
type Uint32 uint32

// Uint32::Cmp
func (a Uint32) Cmp(b_ interface{}) int {
	switch b := b_.(Uint32); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_Uint32 is a cmp.F for type cmp.Uint32
func F_Uint32(a_, b_ interface{}) int {
	switch a, b := a_.(Uint32), b_.(Uint32); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_uint32 is a cmp.F for basic type uint32
func F_uint32(a_, b_ interface{}) int {
	switch a, b := a_.(uint32), b_.(uint32); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// Compare_uint32 is a convenience function for
// explicity comparing basic type uint32
func Compare_uint32(a, b uint32) int {
	switch {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}
