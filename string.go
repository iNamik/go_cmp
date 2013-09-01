package cmp

// String is a cmp.T wrapper for basic type string
type String string

// String::Cmp
func (a String) Cmp(b_ interface{}) int {
	switch b := b_.(String); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_String is a cmp.F for type cmp.String
func F_String(a_, b_ interface{}) int {
	switch a, b := a_.(String), b_.(String); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_string is a cmp.F for basic type string
func F_string(a_, b_ interface{}) int {
	switch a, b := a_.(string), b_.(string); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// Compare_string is a convenience function for
// explicity comparing basic type string
func Compare_string(a, b string) int {
	switch {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}
