package bi

type Data string

func (word *Data) editToken(i int, newToken Data) Data {
	return (*word)[:i] + newToken + (*word)[i+1:]
}
func negate(a Data) Data {
	var c Data
	for i := len(a) - 1; i > 0; i-- {
		c = c + "0"
	}
	c = c + "1"
	b, _ := ADD(c, Data(BitFlip(a, '1')), '0')
	return b
}
