package main

import (
	"fmt"
	tools "simplePL/tools/rpn"
	"strings"
)

func main() {
	str := "(2+((1+2)/4+5)*6/(4+6-4+8-1))*8+2"
	fmt.Println("Input:", str)
	fmt.Println("Output:", tools.ToPostfix(str))
	token := "1 + 2 * 3 *                5"
	tokens := strings.Fields(token)
	for _, value := range tokens {
		fmt.Println(value)
	}

}

// func getDigits(str string) string {
// 	digits, result := "", ""
// 	for _, ch := range str {
// 		if unicode.IsDigit(ch) {
// 			digits += string(ch)
// 			continue
// 		} else if digits != "" {
// 			result += digits + " "
// 			digits = ""
// 		}
// 	}

// 	if digits != "" {
// 		result += digits
// 	}

// 	return result
// }
