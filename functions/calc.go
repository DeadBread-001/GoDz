package functions

import (
	"fmt"
	"github.com/DeadBread-001/GoDz/tree/dz1part2/stack"
	"strconv"
	"strings"
	"unicode"
)

const operators = "+-*/()"

func parseExpression(expression string) ([]string, error) {
	var tokens []string
	var token strings.Builder

	for _, char := range expression {
		stringChar := string(char)

		if stringChar == "." {
			token.WriteString(stringChar)

			continue
		}

		switch {
		case unicode.IsDigit(char):
			token.WriteString(stringChar)
		case strings.Contains(operators, stringChar):
			if token.Len() > 0 {
				tokens = append(tokens, token.String())
				token.Reset()
			}
			tokens = append(tokens, stringChar)
		default:
			return nil, fmt.Errorf("wrong symbol: %v", stringChar)
		}
	}

	if token.Len() > 0 {
		tokens = append(tokens, token.String())
	}

	return tokens, nil
}

func toPolishNotation(expressionTokens []string) ([]string, error) {
	var tokens []string
	var s stack.Stack

	operationPriority := map[rune]int{
		'(': 0,
		'+': 1, '-': 1,
		'*': 2, '/': 2,
		'~': 4,
	}

	for index, token := range expressionTokens {
		if _, err := strconv.ParseFloat(token, 64); err == nil {
			tokens = append(tokens, token)
		}

		switch token {
		case "(":
			s.Push(token)
		case ")":
			for {
				top, isEmpty := s.Top()
				if top == "(" || isEmpty != nil {
					break
				}

				popped, err := s.Pop()
				if err != nil {
					return nil, fmt.Errorf("error in notation conversion: %w", err)
				}
				tokens = append(tokens, popped.(string))
			}

			_, err := s.Pop()
			if err != nil {
				return nil, fmt.Errorf("error in notation conversion: %w", err)
			}
		case "+", "-", "*", "/", "~":
			op := rune(token[0])
			if (op == '-' || op == '+') && (index == 0 || (index > 1 && strings.Contains(operators[:len(operators)-1], expressionTokens[index-1]))) {
				if op == '-' {
					op = '~'
				} else {
					continue
				}
			}

			for {
				top, _ := s.Top()
				if !s.Empty() && operationPriority[rune(top.(string)[0])] >= operationPriority[op] {
					popped, _ := s.Pop()
					tokens = append(tokens, popped.(string))
				} else {
					break
				}
			}
			s.Push(string(op))
		}
	}

	for !s.Empty() {
		popped, _ := s.Pop()
		tokens = append(tokens, popped.(string))
	}

	return tokens, nil
}

func getResult(tokens []string) (float64, error) {
	var s stack.Stack

	for _, token := range tokens {
		if num, err := strconv.ParseFloat(token, 64); err == nil {
			s.Push(num)
		} else {
			switch token[0] {
			case '~':
				last, err := s.Pop()
				if err != nil {
					return 0, fmt.Errorf("error getting result: %w", err)
				}

				s.Push(evaluate('-', 0, last.(float64)))
			default:
				second, err := s.Pop()
				if err != nil {
					return 0, fmt.Errorf("error getting result: %w", err)
				}

				first, err := s.Pop()
				if err != nil {
					return 0, fmt.Errorf("error getting result: %w", err)
				}

				s.Push(evaluate(token[0], first.(float64), second.(float64)))
			}
		}
	}
	num, err := s.Top()
	if err != nil {
		return 0, fmt.Errorf("error getting result: %w", err)
	}

	result, isOk := num.(float64)
	if !isOk {
		return 0, fmt.Errorf("error getting result: %w", err)
	}

	return result, nil
}

func evaluate(op byte, first, second float64) float64 {
	switch op {
	case '+':
		return first + second
	case '-':
		return first - second
	case '*':
		return first * second
	case '/':
		return first / second
	case '~':
		return -second
	}

	return 0
}

func Calculate(expression string) (float64, error) {
	tokens, err := parseExpression(expression)
	if err != nil {
		return 0, err
	}

	tokens, err = toPolishNotation(tokens)
	if err != nil {
		return 0, err
	}

	return getResult(tokens)
}
