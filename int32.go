package cmp

// Int32 is a cmp.T wrapper for basic type int32
type Int32 int32

// Int32::Cmp
func (a Int32) Cmp(b_ interface{}) int {
	switch b := b_.(Int32); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_Int32 is a cmp.F for type cmp.Int32
func F_Int32(a_, b_ interface{}) int {
	switch a, b := a_.(Int32), b_.(Int32); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_int32 is a cmp.F for basic type int32
func F_int32(a_, b_ interface{}) int {
	switch a, b := a_.(int32), b_.(int32); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// Compare_int32 is a convenience function for
// explicity comparing basic type int32
func Compare_int32(a, b int32) int {
	switch {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}
