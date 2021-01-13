package main

import (
	"testing"
)

func TestDiff(t *testing.T) {
	type args struct {
		S string
		T string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case01", args: args{"aacc", "bbdde"}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Diff(tt.args.S, tt.args.T); got != tt.want {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}
