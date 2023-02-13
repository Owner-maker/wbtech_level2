package main

import (
	"testing"
)

func Test_StringUnpack(t *testing.T) {
	tests := []struct {
		name string
		get  string
		want string
	}{
		{"ok", "a4bc2d5e", "aaaabccddddde"},
		{"without_changes", "abcd", "abcd"},
		{"incorrect", "45", ""},
		{"empty", "", ""},
		{"int_as_string", `qwe\4\5`, "qwe45"},
		{"int_as_string_several_times", `qwe\45`, "qwe44444"},
		{"slash_as_string", `qwe\\5`, `qwe\\\\\`},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := StringUnpack(testCase.get)

			if res != testCase.want {
				t.Errorf("got %s, want %s", res, testCase.want)
				t.Error(err)
			}
		})
	}
}
