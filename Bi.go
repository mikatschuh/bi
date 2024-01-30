package bi

import "strconv"

func DIVIDE(Divident, Divisor Data) Data {
	var result Data
	return result
}

func DeciToBi(Decimal int, n int) Data {
	var Binary Data
	a := 1

	for i := 1; i < n; i++ {
		a *= 2
	}

	for i := 0; i < n; i++ {
		if Decimal-a >= 0 {
			Decimal -= a
			Binary = Binary + "1"
		} else {
			Binary = Binary + "0"
		}
		a /= 2
	}
	return Binary
}
func signedDeciToBi(Decimal string, n int) string {
	var Binary string
	Decimald, _ := strconv.Atoi(Decimal)
	a := 1
	for i := 1; i < n-1; i++ {
		a *= 2
	}
	b := a * -2
	if Decimald < 0 {
		Decimald -= b
		Binary = Binary + "1"
	} else {
		Binary = Binary + "0"
	}
	for i := 1; i < n; i++ {
		if Decimald-a >= 0 {
			Decimald -= a
			Binary = Binary + "1"
		} else {
			Binary = Binary + "0"
		}
		a /= 2
	}
	return Binary
}
func BoolToBit(condition bool) Bit {
	if condition {
		return '1'
	} else {
		return '0'
	}
}
func ADD(a Data, b Data, c Bit) (Data, Bit) {
	var result Data
	for i := len(a) - 1; i >= 0; i-- {
		va := Bit(a[i])
		bv := Bit(b[i])
		axorb := XOR(va, bv)
		result = result + XOR(c, axorb).Data()
		c = OR(AND(va, bv), AND(axorb, c))
	}
	return result, c
}
func AND(a Bit, b Bit) Bit {
	if a.bool() && b.bool() {
		return '1'
	}
	return '0'
}
func OR(a Bit, b Bit) Bit {
	if a.bool() || b.bool() {
		return '1'
	}
	return '0'
}
func NOT(a Bit) Bit {
	if a.bool() {
		return '1'
	}
	return '0'
}
func NAND(a Bit, b Bit) Bit {
	if !(a.bool() && b.bool()) {
		return '1'
	}
	return '0'
}
func NOR(a Bit, b Bit) Bit {
	if !(a.bool() || b.bool()) {
		return '1'
	}
	return '0'
}
func XOR(a Bit, b Bit) Bit {
	if a != b {
		return '1'
	}
	return '0'
}
func XNOR(a Bit, b Bit) Bit {
	if a == b {
		return '1'
	}
	return '0'
}
