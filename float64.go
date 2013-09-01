package cmp

// Float64 is a cmp.T wrapper for basic type float64
type Float64 float64

// Float64::Cmp
func (a Float64) Cmp(b_ interface{}) int {
	switch b := b_.(Float64); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_Float64 is a cmp.F for type cmp.Float64
func F_Float64(a_, b_ interface{}) int {
	switch a, b := a_.(Float64), b_.(Float64); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_float64 is a cmp.F for basic type float64
func F_float64(a_, b_ interface{}) int {
	switch a, b := a_.(float64), b_.(float64); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// Compare_float64 is a convenience function for
// explicity comparing basic type float64
func Compare_float64(a, b float64) int {
	switch {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}
