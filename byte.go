package cmp

// Byte is a cmp.T wrapper for basic type byte
type Byte byte

// Byte::Cmp
func (a Byte) Cmp(b_ interface{}) int {
	switch b := b_.(Byte); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_Byte is a cmp.F for type cmp.Byte
func F_Byte(a_, b_ interface{}) int {
	switch a, b := a_.(Byte), b_.(Byte); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// F_byte is a cmp.F for basic type byte
func F_byte(a_, b_ interface{}) int {
	switch a, b := a_.(byte), b_.(byte); {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}

// Compare_byte is a convenience function for
// explicity comparing basic type byte
func Compare_byte(a, b byte) int {
	switch {
	case a < b:
		return LT
	case a > b:
		return GT
	default:
		return EQ
	}
}
