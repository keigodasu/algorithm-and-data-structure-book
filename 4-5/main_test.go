package main

import "testing"

func TestFibo(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case01",
			args: args{N: 3},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibo(tt.args.N); got != tt.want {
				t.Errorf("Fibo() = %v, want %v", got, tt.want)
			}
		})
	}
}
