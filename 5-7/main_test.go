package main

import "testing"

func TestKnapsack(t *testing.T) {
	type args struct {
		weight []int
		value  []int
		W      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case01",
			args{
				weight: []int{2, 1, 3, 2, 1, 5},
				value:  []int{3, 2, 6, 1, 3, 85},
				W:      9,
			},
			94,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Knapsack(tt.args.weight, tt.args.value, tt.args.W); got != tt.want {
				t.Errorf("Knapsack() = %v, want %v", got, tt.want)
			}
		})
	}
}
