package cmp

// Int is a cmp.T wrapper for basic type int
type Int64 int64

// Int64::Cmp
func (a Int64) Cmp(b_ interface{}) int {
	switch b := b_.(Int64); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_Int64 is a cmp.F for type cmp.Int64
func F_Int64(a_, b_ interface{}) int {
	switch a, b := a_.(Int64), b_.(Int64); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_int64 is a cmp.F for basic type int64
func F_int64(a_, b_ interface{}) int {
	switch a, b := a_.(int64), b_.(int64); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// Compare_int64 is a convenience function for
// explicity comparing basic type int64
func Compare_int64(a, b int64) int {
	switch {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}
