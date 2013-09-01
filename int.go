package cmp

// Int is a cmp.T wrapper for basic type int
type Int int

// Int::Cmp
func (a Int) Cmp(b_ interface{}) int {
	switch b := b_.(Int); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_Int is a cmp.F for type cmp.Int
func F_Int(a_, b_ interface{}) int {
	switch a, b := a_.(Int), b_.(Int); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_int is a cmp.F for basic type int
func F_int(a_, b_ interface{}) int {
	switch a, b := a_.(int), b_.(int); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// Compare_int is a convenience function for
// explicity comparing basic type int
func Compare_int(a, b int) int {
	switch {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}
