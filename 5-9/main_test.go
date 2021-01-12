package main

import "testing"

func TestSeparateCost(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{10},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run("case", func(t *testing.T) {
			if got := SeparateCost(tt.args.N); got != tt.want {
				t.Errorf("SeparateCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
