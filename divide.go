package main

import "fmt"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if dividend == divisor {
		return 1
	}
	result := 0
	flag := (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0)
	dividend = abs(dividend)
	divisor = abs(divisor)
	for divisor <= dividend {
		if dividend >= divisor && dividend < divisor+divisor {
			// 介于1倍-2倍之间就是只需要次数+1就行了，然后应该就退出了去了
			result++
			dividend -= divisor
		} else {
			// 每次都x2，直到超过了被除数，然后次数和被除数向下减
			d, t := divisor, 1
			for d < dividend {
				d = d << 1
				t = t << 1
			}
			result += t >> 1
			dividend -= d >> 1
		}
	}
	if flag {
		result = -result
	}
	if result < -2147483648 {
		return -2147483648
	}
	if result > 2147483647 {
		return 2147483647
	}
	return result
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	// fmt.Println(divide(10, 3))
	// fmt.Println(divide(7, -3))
	fmt.Println(divide(1, 1))
}
