package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	list := []int{}
	for _, v := range tokens {
		num, err := strconv.Atoi(v)
		if err == nil {
			list = append(list, num)
			continue
		}
		num1 := list[len(list)-1]
		num2 := list[len(list)-2]
		res := 0
		switch v {
		case "+":
			res = num2 + num1
			break
		case "-":
			res = num2 - num1
			break
		case "*":
			res = num2 * num1
			break
		case "/":
			res = num2 / num1
			break
		default:
			continue
		}
		list = list[:len(list)-2]
		list = append(list, res)
	}
	return list[0]
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
}
