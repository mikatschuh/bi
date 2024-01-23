package bi

type Bit byte

func (b Bit) Bool() bool {
	if b == '1' {
		return true
	} else {
		return false
	}
}
func (b *Bit) NOT() {
	if (*b) == '1' {
		*b = '0'
	} else {
		*b = '1'
	}
}
