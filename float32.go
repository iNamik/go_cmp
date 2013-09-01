package cmp

// Float32 is a cmp.T wrapper for basic type float32
type Float32 float32

// Float32::Cmp
func (a Float32) Cmp(b_ interface{}) int {
	switch b := b_.(Float32); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_Float32 is a cmp.F for type cmp.Float32
func F_Float32(a_, b_ interface{}) int {
	switch a, b := a_.(Float32), b_.(Float32); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_float32 is a cmp.F for basic type float32
func F_float32(a_, b_ interface{}) int {
	switch a, b := a_.(float32), b_.(float32); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// Compare_float32 is a convenience function for
// explicity comparing basic type float32
func Compare_float32(a, b float32) int {
	switch {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}
