package cmp

// Int16 is a cmp.T wrapper for basic type int16
type Int16 int16

// Int16::Cmp
func (a Int16) Cmp(b_ interface{}) int {
	switch b := b_.(Int16); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_Int16 is a cmp.F for type cmp.Int16
func F_Int16(a_, b_ interface{}) int {
	switch a, b := a_.(Int16), b_.(Int16); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_int16 is a cmp.F for basic type int16
func F_int16(a_, b_ interface{}) int {
	switch a, b := a_.(int16), b_.(int16); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// Compare_int16 is a convenience function for
// explicity comparing basic type int16
func Compare_int16(a, b int16) int {
	switch {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}
