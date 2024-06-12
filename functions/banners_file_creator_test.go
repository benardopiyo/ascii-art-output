package functions

import (
	"reflect"
	"testing"
)

func TestBannerArt(t *testing.T) {
	testCases := []struct {
		name     string
		args     string
		expected map[int][]string
	}{
		{name: "standard", args: "../banners/standard.txt", expected: BannerArt("../banners/standard.txt")},
		{name: "shadow", args: "../banners/shadow.txt", expected: BannerArt("../banners/shadow.txt")},
		{name: "thinkertoy", args: "../banners/thinkertoy.txt", expected: BannerArt("../banners/thinkertoy.txt")},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := BannerArt(tc.args); !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("BannerArt() = %v, want %v", got, tc.expected)
			}
		})
	}
}
