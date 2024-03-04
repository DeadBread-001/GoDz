package uniq_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DeadBread-001/GoDz/tree/dz1part1/uniq"
)

var errNum = errors.New("invalid input number")

func TestIgnoreCase(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"Hello", "WORLD", "Go"}, []string{"hello", "world", "go"}},
		{[]string{"One", "two", "Three"}, []string{"one", "two", "three"}},
	}

	for _, test := range tests {
		result := uniq.IgnoreCase(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("IgnoreCase(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestIgnoreFields(t *testing.T) {
	tests := []struct {
		input    []string
		num      int
		expected []string
		err      error
	}{
		{[]string{"a b c", "1 2 3", "x y z"}, 1, []string{"b c", "2 3", "y z"}, nil},
		{[]string{"word1 word2 word3", "a b c"}, 3, []string{"", ""}, nil},
		{[]string{"a", "b", "c"}, -1, nil, errNum},
	}

	for _, test := range tests {
		result, err := uniq.IgnoreFields(test.input, test.num)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("IgnoreFields(%v, %d) = %v, expected %v", test.input, test.num, result, test.expected)
		}

		if !reflect.DeepEqual(err, test.err) {
			t.Errorf("IgnoreFields(%v, %d) returned error %v, expected %v", test.input, test.num, err, test.err)
		}
	}
}

func TestIgnoreCharacters(t *testing.T) {
	tests := []struct {
		input    []string
		num      int
		expected []string
		err      error
	}{
		{[]string{"abcdef", "12345", "xyz"}, 2, []string{"cdef", "345", "z"}, nil},
		{[]string{"word1", "word2", "word3"}, 4, []string{"1", "2", "3"}, nil},
		{[]string{"a", "b", "c"}, -1, nil, errNum},
	}

	for _, test := range tests {
		result, err := uniq.IgnoreCharacters(test.input, test.num)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("IgnoreCharacters(%v, %d) = %v, expected %v", test.input, test.num, result, test.expected)
		}

		if !reflect.DeepEqual(err, test.err) {
			t.Errorf("IgnoreCharacters(%v, %d) returned error %v, expected %v", test.input, test.num, err, test.err)
		}
	}
}
