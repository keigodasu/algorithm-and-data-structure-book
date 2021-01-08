package main

import "testing"

func TestGcd(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "found",
			args: args{
				m: 15,
				n: 6,
			},
			want: 3,
		},
		{
			name: "found 1",
			args: args{
				m: 15,
				n: 7,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gcd(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("Gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}
