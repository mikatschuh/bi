package bi

type Data string

func (word *Data) editToken(i int, newToken Data) Data {
	return (*word)[:i] + newToken + (*word)[i+1:]
}
func (a Data) negate() Data {
	var c Data
	for i := len(a) - 1; i > 0; i-- {
		c = c + "0"
	}
	c = c + "1"
	b, _ := ADD(c, Data(BitFlip(a, '1')), '0')
	return b
}
func (a *Data) NegateIt() Data {
	var c Data
	for i := len(*a) - 1; i > 0; i-- {
		c = c + "0"
	}
	c = c + "1"
	b, _ := ADD(c, Data(BitFlip(*a, '1')), '0')
	return b
}
func BiToDeci(Binary Data) int {
	n := 1
	Decimal := 0

	for i := len(Binary) - 1; i >= 0; i-- {
		if Binary[i] == '1' {
			Decimal += n
		}
		n *= 2
	}
	return Decimal
}
