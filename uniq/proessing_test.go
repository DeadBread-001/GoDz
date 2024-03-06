package uniq

import (
	"reflect"
	"testing"
)

func TestExecuteUniq(t *testing.T) {
	tests := []struct {
		name   string
		flags  Options
		input  []string
		output []string
		err    error
	}{
		{
			name:   "No flags",
			flags:  Options{},
			input:  []string{"Apple", "apple", "apple", "orange", "banana", "grape"},
			output: []string{"Apple", "apple", "orange", "banana", "grape"},
			err:    nil,
		},
		{
			name:   "IgnoreCase flag",
			flags:  Options{iFlag: true},
			input:  []string{"apple", "APPLE", "Banana", "Orange", "Grape", "GraPe"},
			output: []string{"apple", "Banana", "Orange", "Grape"},
			err:    nil,
		},
		{
			name:   "IgnoreFields flag",
			flags:  Options{fFlag: 1},
			input:  []string{"apple orange", "banana orange", "orange lemon", "orange pear", "banana cherry", "grape cherry"},
			output: []string{"apple orange", "orange lemon", "orange pear", "banana cherry"},
			err:    nil,
		},
		{
			name:   "IgnoreCharacters flag",
			flags:  Options{sFlag: 3},
			input:  []string{"ApPle", "apple", "appde", "orange", "banana", "graana"},
			output: []string{"ApPle", "appde", "orange", "banana"},
			err:    nil,
		},
		{
			name:   "CountOfLines flag",
			flags:  Options{cFlag: true},
			input:  []string{"apple", "apple", "banana", "apple", "grape", "grape", "grape"},
			output: []string{"2 apple", "1 banana", "1 apple", "3 grape"},
			err:    nil,
		},
		{
			name:   "DuplicatedLines flag",
			flags:  Options{dFlag: true},
			input:  []string{"apple", "apple", "banana", "apple", "orange", "orange", "grape"},
			output: []string{"apple", "orange"},
			err:    nil,
		},
		{
			name:   "UniqLines flag",
			flags:  Options{uFlag: true},
			input:  []string{"apple", "apple", "banana", "apple", "orange", "orange", "grape"},
			output: []string{"banana", "apple", "grape"},
			err:    nil,
		},
		{
			name:   "IgnoreCase + IgnoreFields",
			flags:  Options{iFlag: true, fFlag: 1},
			input:  []string{"apple orange", "APPLE ORANGE", "APPLE lemon", "biba lemon", "banana aboba", "bimba aBOBa"},
			output: []string{"apple orange", "APPLE lemon", "banana aboba"},
			err:    nil,
		},
		{
			name:   "IgnoreCharacters + CountOfLines",
			flags:  Options{sFlag: 3, cFlag: true},
			input:  []string{"ApPle", "apple", "appde", "orange", "banana", "graana"},
			output: []string{"2 apple", "1 appde", "1 orange", "2 graana"},
			err:    nil,
		},
		{
			name:   "Invalid flag combination",
			flags:  Options{cFlag: true, dFlag: true},
			input:  []string{"apple", "banana", "apple", "orange", "banana", "grape"},
			output: nil,
			err:    errParsing,
		},
		{
			name:   "DuplicatedLines + UniqLines",
			flags:  Options{dFlag: true, uFlag: true},
			input:  []string{"apple", "banana", "apple", "orange", "banana", "grape"},
			output: nil,
			err:    errParsing,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ExecuteUniq(test.flags, test.input)

			if !reflect.DeepEqual(result, test.output) {
				t.Errorf("ExecuteUniq(%v, %v) = %v, expected %v", test.flags, test.input, result, test.output)
			}

			if !reflect.DeepEqual(err, test.err) {
				t.Errorf("ExecuteUniq(%v, %v) returned error %v, expected %v", test.flags, test.input, err, test.err)
			}
		})
	}
}
