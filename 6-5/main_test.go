package main

import "testing"

func Test_binarysearch(t *testing.T) {
	type args struct {
		h []int
		s []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"caes01",
			args{
				h: []int{1, 2, 3},
				s: []int{1, 1, 1},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarysearch(tt.args.h, tt.args.s); got != tt.want {
				t.Errorf("binarysearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
