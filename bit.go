package bi

type Bit byte

func (b Bit) bool() bool {
	if b == '1' {
		return true
	} else {
		return false
	}
}
func (b Bit) not() Bit {
	if (b) == '1' {
		return '0'
	} else {
		return '1'
	}
}
func (b Bit) boolNot() bool {
	if (b) == '1' {
		b = '0'
	} else {
		b = '1'
	}
	if b == '1' {
		return true
	} else {
		return false
	}
}
func (b *Bit) Flip() {
	if (*b) == '1' {
		*b = '0'
	} else {
		*b = '1'
	}
}
func (b Bit) Data() Data {
	return Data(b)
}
