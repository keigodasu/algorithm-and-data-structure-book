package main

import "testing"

func TestDpFrog(t *testing.T) {
	type args struct {
		h []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case01",
			args: args{h: []int{2, 9, 4, 5, 1, 6, 10}},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DpFrog(tt.args.h); got != tt.want {
				t.Errorf("DpFrog() = %v, want %v", got, tt.want)
			}
		})
	}
}
