package ansi

import (
	"testing"
)

func TestNewStyle(t *testing.T) {
	tests := []struct {
		inpCodes     []string
		expectedArgs string
		expectedCode string
	}{
		{[]string{}, "0", "\x1b[0m"},
		{[]string{"0"}, "0", "\x1b[0m"},
		{[]string{"8"}, "8", "\x1b[8m"},
		{[]string{"23"}, "23", "\x1b[23m"},
		{[]string{"1", "2"}, "1;2", "\x1b[1;2m"},
		{[]string{"2", "3"}, "2;3", "\x1b[2;3m"},
		{[]string{"3", "4"}, "3;4", "\x1b[3;4m"},
		{[]string{"4", "5"}, "4;5", "\x1b[4;5m"},
		{[]string{"5", "6"}, "5;6", "\x1b[5;6m"},
		{[]string{"6", "7"}, "6;7", "\x1b[6;7m"},
		{[]string{"7", "8"}, "7;8", "\x1b[7;8m"},
		{[]string{"1", "2", "3"}, "1;2;3", "\x1b[1;2;3m"},
		{[]string{"7", "3", "2", "4"}, "7;3;2;4", "\x1b[7;3;2;4m"},
		{[]string{"10", "7", "3", "2", "4"}, "10;7;3;2;4", "\x1b[10;7;3;2;4m"},
	}

	for _, tt := range tests {
		style := NewStyle(tt.inpCodes...)

		if style.Args() != tt.expectedArgs {
			t.Errorf("Unexpected value for Style.args, want=%s, got=%s",
				tt.expectedArgs, style.Args())
		}

		if style.StyleCode() != tt.expectedCode {
			t.Errorf("Unexpected value for Style.StyleCode, want=%q, got=%q",
				tt.expectedCode, style.StyleCode())
		}

	}
}

func TestLastChar(t *testing.T) {
	tests := []struct {
		input    string
		expected byte
	}{
		{input: "", expected: 0},
		{input: "hello", expected: 'o'},
		{input: "hello\n", expected: '\n'},
	}

	for _, tt := range tests {
		res := lastChar(tt.input)

		if res != tt.expected {
			t.Errorf("Unexpected result for lastChar, want=%q, got=%q",
				tt.expected, res)
		}
	}
}
