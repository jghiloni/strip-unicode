package unistripper_test

import (
	"context"
	"strings"
	"testing"

	"github.com/jghiloni/strip-unicode/unistripper"
)

var testCases = []struct {
	inputString  string
	outputString string
}{
	{
		"no unicode",
		"no unicode",
	},
	{
		"hidden character\u200b",
		"hidden character",
	},
	{
		"with emoji 🤯",
		"with emoji ",
	},
	{
		"with ∑ character",
		"with  character",
	},
	{
		"multi\nline",
		"multi\nline",
	},
}

func TestStripUnicode(t *testing.T) {
	for _, tt := range testCases {
		output := new(strings.Builder)
		err := unistripper.StripUnicode(context.Background(), strings.NewReader(tt.inputString), output)
		if err != nil {
			t.Error("unexpected error", err)
		}

		if tt.outputString != output.String() {
			t.Errorf("actual output string %q does not match expected output string %q", output.String(), tt.outputString)
		}
	}
}
