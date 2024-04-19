package tools

import (
	"simplePL/tools"
	"strings"
	"unicode"
)

func ToPostfix(str string) string {
	operators := tools.Stack{}
	exitString := ""
	digits := ""
	for i, ch := range str {
		if ch == ' ' {
			continue
		}

		if unicode.IsDigit(ch) {
			digits += string(ch)
			if i == len(str)-1 {
				exitString += digits + " "
				digits = ""
			}
			continue
		}

		if digits != "" {
			exitString += digits + " "
			digits = ""
		}

		if isOperator(ch) {
			for !operators.IsEmpty() {
				top, _ := operators.Peek()
				if preference(ch) <= preference(top.(rune)) && top != '(' {
					top, _ := operators.Pop()
					topStr := string(top.(rune))
					exitString += topStr + " "
				} else {
					break
				}
			}
			operators.Push(ch)
			continue
		}

		if ch == '(' {
			operators.Push(ch)
			continue
		}

		if ch == ')' {
			for {
				top, err := operators.Pop()
				if err != nil || top == '(' {
					break
				}
				exitString += string(top.(rune)) + " "
			}
			continue
		}

	}

	for !operators.IsEmpty() {
		top, _ := operators.Pop()
		topStr := string(top.(rune))
		exitString += topStr + " "
	}

	if len(digits) != 0 {
		exitString += digits
	}
	return exitString
}

// func Calculate(str string) int {
// 	tokens := strings.Fields(str)
// 	stack := tools.Stack{}
// 	for _, ch := range tokens {
// 		num, err := strconv.Atoi(ch)
// 		if err == nil {
// 			stack.Push(num)
// 		}

// 		op := isOperator()

// 	}
// }

func isOperator(ch rune) bool {
	operator := `+-*/^`
	return strings.ContainsRune(operator, ch)
}

func preference(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	case '^':
		return 3
	default:
		return 0
	}
}
