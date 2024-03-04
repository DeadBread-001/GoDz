package uniq_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DeadBread-001/GoDz/tree/dz1part1/uniq"
)

var errProcess = errors.New("selection funcs error")

func TestBuildResult(t *testing.T) {
	tests := []struct {
		suitableLines []bool
		lines         []string
		expected      []string
		err           error
	}{
		{[]bool{true, false, true}, []string{"one", "two", "three"}, []string{"one", "three"}, nil},
		{[]bool{true, true, true}, []string{"a", "b", "c"}, []string{"a", "b", "c"}, nil},
		{[]bool{true, false}, []string{"x", "y", "z"}, nil, errProcess},
	}

	for _, test := range tests {
		result, err := uniq.BuildResult(test.suitableLines, test.lines)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("BuildResult(%v, %v) = %v, expected %v", test.suitableLines, test.lines, result, test.expected)
		}

		if !reflect.DeepEqual(err, test.err) {
			t.Errorf("BuildResult(%v, %v) returned error %v, expected %v", test.suitableLines, test.lines, err, test.err)
		}
	}
}

func TestStandartUniq(t *testing.T) {
	tests := []struct {
		lines    []string
		expected []bool
	}{
		{[]string{"a", "a", "b", "c", "c", "c"}, []bool{true, false, true, true, false, false}},
		{[]string{"one", "two", "two", "three"}, []bool{true, true, false, true}},
		{[]string{"x", "y", "z"}, []bool{true, true, true}},
	}

	for _, test := range tests {
		result := uniq.StandartUniq(test.lines)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("StandartUniq(%v) = %v, expected %v", test.lines, result, test.expected)
		}
	}
}

func TestDuplicatedLines(t *testing.T) {
	tests := []struct {
		lines    []string
		expected []bool
	}{
		{[]string{"a", "a", "b", "c", "c", "c"}, []bool{true, false, false, true, false, false}},
		{[]string{"one", "two", "two", "three"}, []bool{false, true, false, false}},
		{[]string{"x", "y", "z"}, []bool{false, false, false}},
	}

	for _, test := range tests {
		result := uniq.DuplicatedLines(test.lines)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("DuplicatedLines(%v) = %v, expected %v", test.lines, result, test.expected)
		}
	}
}

func TestUniqLines(t *testing.T) {
	tests := []struct {
		lines    []string
		expected []bool
	}{
		{[]string{"a", "a", "b", "c", "c", "c"}, []bool{false, false, true, false, false, false}},
		{[]string{"one", "two", "two", "three"}, []bool{true, false, false, true}},
		{[]string{"x", "y", "z"}, []bool{true, true, true}},
	}

	for _, test := range tests {
		result := uniq.UniqLines(test.lines)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("UniqLines(%v) = %v, expected %v", test.lines, result, test.expected)
		}
	}
}

func TestCountOfLines(t *testing.T) {
	tests := []struct {
		lines         []string
		originalLines []string
		expected      []bool
	}{
		{[]string{"a", "a", "b", "c", "c", "c"}, []string{"2 a", "1 b", "3 c", "", "", ""}, []bool{false, true, true, false, false, true}},
		{[]string{"one", "two", "two", "three"}, []string{"1 one", "2 two", "1 three", ""}, []bool{true, false, true, true}},
		{[]string{"x", "y", "z"}, []string{"1 x", "1 y", "1 z"}, []bool{true, true, true}},
	}

	for _, test := range tests {
		result := uniq.CountOfLines(test.lines, test.originalLines)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("CountOfLines(%v, %v) = %v, expected %v", test.lines, test.originalLines, result, test.expected)
		}
	}
}
