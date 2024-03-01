package functions

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestParseExpression(t *testing.T) {
	tests := []struct {
		expression string
		expected   []string
		err        error
	}{
		{"3+4*2/(1-5)", []string{"3", "+", "4", "*", "2", "/", "(", "1", "-", "5", ")"}, nil},
		{"2.5+3.7", []string{"2.5", "+", "3.7"}, nil},
		{"invalid expression", nil, fmt.Errorf("wrong symbol: i")},
	}

	for _, test := range tests {
		result, err := parseExpression(test.expression)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("parseExpression(%v) = %v, expected %v", test.expression, result, test.expected)
		}

		if !reflect.DeepEqual(err, test.err) {
			t.Errorf("parseExpression(%v) returned error %v, expected %v", test.expression, err, test.err)
		}
	}
}

func TestToPolishNotation(t *testing.T) {
	tests := []struct {
		expressionTokens []string
		expected         []string
		err              error
	}{
		{[]string{"3", "+", "4", "*", "2", "/", "(", "1", "-", "5", ")"}, []string{"3", "4", "2", "*", "1", "5", "-", "/", "+"}, nil},
		{[]string{"2.5", "+", "3.7"}, []string{"2.5", "3.7", "+"}, nil},
		{[]string{"3", "+", "&", "*", "expression"}, []string{"3", "*", "+"}, nil},
	}

	for _, test := range tests {
		result, err := toPolishNotation(test.expressionTokens)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("toPolishNotation(%v) = %v, expected %v", test.expressionTokens, result, test.expected)
		}

		if !reflect.DeepEqual(err, test.err) {
			t.Errorf("toPolishNotation(%v) returned error %v, expected %v", test.expressionTokens, err, test.err)
		}
	}
}

func TestGetResult(t *testing.T) {
	tests := []struct {
		tokens   []string
		expected float64
		err      error
	}{
		{[]string{"3", "4", "2", "*", "1", "5", "-", "/", "+"}, 1, nil},
		{[]string{"2.5", "3.7", "+"}, 6.2, nil},
		{[]string{"3", "4", "invalid", "*", "expression"}, 0, fmt.Errorf("error getting result: tried to pop from empty stack")},
	}

	for _, test := range tests {
		result, err := getResult(test.tokens)

		if math.Abs(result-test.expected) > 1e-9 {
			t.Errorf("getResult(%v) = %v, expected %v", test.tokens, result, test.expected)
		}

		if err != nil && test.err != nil {
			if err.Error() != test.err.Error() {
				t.Errorf("getResult(%v) returned error %v, expected %v", test.tokens, err, test.err)
			}
		} else if !errors.Is(err, test.err) {
			t.Errorf("getResult(%v) returned error %v, expected %v", test.tokens, err, test.err)
		}
	}
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		expression string
		expected   float64
		err        error
	}{
		{"3+4*2/(1-5)", 1, nil},
		{"2.5+3.7", 6.2, nil},
		{"invalid expression", 0, fmt.Errorf("wrong symbol: i")},
	}

	for _, test := range tests {
		result, err := Calculate(test.expression)

		if math.Abs(result-test.expected) > 1e-9 {
			t.Errorf("Calculate(%v) = %v, expected %v", test.expression, result, test.expected)
		}

		if !reflect.DeepEqual(err, test.err) {
			t.Errorf("Calculate(%v) returned error %v, expected %v", test.expression, err, test.err)
		}
	}
}
