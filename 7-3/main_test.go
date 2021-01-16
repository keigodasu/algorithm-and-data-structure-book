package main

import "testing"

func TestMultipleArray(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case01",
			args: args{
				A: []int{1, 1, 1},
				B: []int{2, 2, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultipleArray(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("MultipleArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
