package cmp

// Int8 is a cmp.T wrapper for basic type int8
type Int8 int8

// Int8::Cmp
func (a Int8) Cmp(b_ interface{}) int {
	switch b := b_.(Int8); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_Int8 is a cmp.F for type cmp.Int8
func F_Int8(a_, b_ interface{}) int {
	switch a, b := a_.(Int8), b_.(Int8); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_int8 is a cmp.F for basic type int8
func F_int8(a_, b_ interface{}) int {
	switch a, b := a_.(int8), b_.(int8); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// Compare_int8 is a convenience function for
// explicity comparing basic type int8
func Compare_int8(a, b int8) int {
	switch {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}
