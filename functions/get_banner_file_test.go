package functions

import (
	"os"
	"testing"
)

func TestFileName(t *testing.T) {
	testCases := []struct {
		name     string
		args     []string
		expected string
	}{
		{"NoArgs", []string{"progName"}, ""},
		{"TwoArgs", []string{"progName", "arg1"}, "standard.txt"},
		{"ThreeArgs", []string{"progName", "arg1", "arg2"}, "standard.txt"},
		{"FourArgsStandard", []string{"progName", "arg1", "arg2", "standard"}, "standard.txt"},
		{"FourArgsThinkertoy", []string{"progName", "arg1", "arg2", "thinkertoy"}, "thinkertoy.txt"},
		{"FourArgsShadow", []string{"progName", "arg1", "arg2", "shadow"}, "shadow.txt"},
		{"FourArgsInvalid", []string{"progName", "arg1", "arg2", "invalid"}, "Invalid bannerfile name"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			originalArgs := os.Args                   // Save the original os.Args
			defer func() { os.Args = originalArgs }() // Restore os.Args after the test
			os.Args = tc.args                         // Set os.Args to the test case arguments

			if got := FileName(); got != tc.expected {
				t.Errorf("FileName() = %v, want %v", got, tc.expected)
			}
		})
	}
}
