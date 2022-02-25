package utils

import (
	"fmt"
	"testing"
)

func TestDirectionComplement(t *testing.T) {
	var tests = []struct {
		given *Direction
		want  *Direction
	}{
		{NORTH, SOUTH},
		{SOUTH, NORTH},
		{EAST, WEST},
		{WEST, EAST},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s Complement", tt.given.name)
		t.Run(testname, func(t *testing.T) {
			ans, _ := tt.given.Complement()
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans.name, tt.want.name)
			}
		})
	}
}
