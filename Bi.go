package bi

import "strconv"

func Dividing(Divident, Divisor Data) Data {
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
func BitFlip(a Data, enable byte) Data {
	var result Data
	if enable == '1' {
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
func ADD(a Data, b Data, c byte) (Data, byte) {
	result := a
	return result, c
}
func signedBiToDeci(Binary Data) int {
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
