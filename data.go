package bi

type Data string

func (word *Data) EditToken(i int, newToken Data) Data {
	return (*word)[:i] + newToken + (*word)[i+1:]
}

// returns Negation
func (a Data) negate() Data {
	var c Data
	for i := len(a) - 1; i > 0; i-- {
		c = c + "0"
	}
	c = c + "1"
	b, _ := ADD(c, Data(a.flip()), '0')
	return b
}

// makes binary numbers negative
func (a *Data) NegateIt() Data {
	var c Data
	for i := len(*a) - 1; i > 0; i-- {
		c = c + "0"
	}
	c = c + "1"
	b, _ := ADD(c, Data(a.flip()), '0')
	return b
}

// translation from Bi to Decimal. Can be used like that: binary.BiToDeci()
func (Binary Data) deci() int {
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

// translates Binary into a signed integer
func (Binary Data) signedDeci() int {
	n := 1
	Decimal := 0
	for i := 1; i < len(Binary)-1; i++ {
		n *= 2
	}
	if Binary[0] == '1' {
		Decimal -= n * 2
	}
	for i := 1; i < len(Binary); i++ {
		if Binary[i] == '1' {
			Decimal += n
		}
		n /= 2
	}
	return Decimal
}

// filps a if enable is on
func (a *Data) FlipIf(enable bool) {
	var result Data
	if enable {
		for _, value := range *a {
			if value == '1' {
				result = result + "0"
			} else {
				result = result + "1"
			}
		}
	} else {
		result = *a
	}
	*a = result
}
func (a *Data) Flip() {
	var result Data
	for _, value := range *a {
		if value == '1' {
			result = result + "0"
		} else {
			result = result + "1"
		}
	}
	*a = result
}
func (a Data) flipIf(enable bool) Data {
	var result Data
	if enable {
		for _, value := range a {
			if value == '1' {
				result = result + "0"
			} else {
				result = result + "1"
			}
		}
	} else {
		result = a
	}
	return result
}
func (a Data) flip() Data {
	var result Data
	for _, value := range a {
		if value == '1' {
			result = result + "0"
		} else {
			result = result + "1"
		}
	}
	return result
}
