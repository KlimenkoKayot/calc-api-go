package calculation

import (
	"strconv"
	"unicode"
)

type Node struct {
	Value     interface{}
	IsFloat64 bool
}

func ExpressionToStack(expresison string) ([]Node, error) {
	result := []Node{}
	runes := []rune(expresison)
	for j := 0; j < len(runes); j++ {
		switch runes[j] {
		case rune(' '):
			continue
		case rune('+'):
			result = append(result, Node{Value: string('+'), IsFloat64: false})
		case rune('-'):
			result = append(result, Node{Value: string('-'), IsFloat64: false})
		case rune('*'):
			result = append(result, Node{Value: string('*'), IsFloat64: false})
		case rune('/'):
			result = append(result, Node{Value: string('/'), IsFloat64: false})
		case rune('('):
			result = append(result, Node{Value: string('('), IsFloat64: false})
		case rune(')'):
			result = append(result, Node{Value: string(')'), IsFloat64: false})
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0': // начинается число
			var num float64 = 0
			lastIdx := -1
			for i := j + 1; i < len(runes)+1; i++ {
				numTmp, err := strconv.ParseFloat(string(runes[j:i]), 64)
				if err == nil {
					lastIdx = i
					num = numTmp
				}
				if i == len(runes) || (string(runes[i]) != "." && !unicode.IsDigit(runes[i])) {
					break
				}
			}
			result = append(result, Node{Value: num, IsFloat64: true})
			j = lastIdx - 1
		default:
			return nil, ErrInvalidSymbolExpression
		}
	}
	return result, nil
}

func ExpressionToRPN(expresison string) ([]interface{}, error) {
	all, err := ExpressionToStack(expresison)
	if err != nil {
		return nil, err
	}

	priority := map[interface{}]int{
		"(": 0,
		")": 1,
		"+": 2,
		"-": 2,
		"*": 3,
		"/": 3,
	}
	result := []interface{}{}
	stack := []interface{}{}
	for _, val := range all {
		if val.IsFloat64 {
			result = append(result, val.Value)
		} else {
			if len(stack) == 0 {
				stack = append(stack, val.Value)
			} else {
				if val.Value == "(" {
					stack = append(stack, "(")
					continue
				}
				if val.Value != ")" {
					for len(stack) > 0 && priority[stack[len(stack)-1]] >= priority[val.Value] {
						result = append(result, stack[len(stack)-1])
						stack = stack[:len(stack)-1]
					}
					stack = append(stack, val.Value)
				} else {
					for {
						if len(stack) == 0 {
							return nil, ErrBadReversedPoland
						}
						if stack[len(stack)-1] == "(" {
							stack = stack[:len(stack)-1]
							break
						}
						result = append(result, stack[len(stack)-1])
						stack = stack[:len(stack)-1]
					}
				}
			}
		}
	}
	for len(stack) > 0 {
		result = append(result, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return result, nil
}

func Calc(expression string) (float64, error) {
	values, err := ExpressionToRPN(expression)
	if err != nil {
		return 0.0, err
	}
	val := []float64{}
	for i := 0; i < len(values); i++ {
		if num, ok := values[i].(float64); ok {
			val = append(val, num)
		} else {
			if len(val) < 2 {
				return 0.0, ErrInvalidExpression
			}

			b := val[len(val)-1]
			val = val[:len(val)-1]
			a := val[len(val)-1]
			val = val[:len(val)-1]

			switch values[i].(string) {
			case "+":
				val = append(val, a+b)
			case "-":
				val = append(val, a-b)
			case "*":
				val = append(val, a*b)
			case "/":
				if b == 0 {
					return 0.0, ErrDevisionByZero
				}
				val = append(val, a/b)
			}
		}
	}
	if len(val) > 1 {
		return 0.0, ErrInvalidExpression
	}
	if len(val) == 0 {
		return 0.0, ErrEmptyExpression
	}
	return val[0], nil
}
