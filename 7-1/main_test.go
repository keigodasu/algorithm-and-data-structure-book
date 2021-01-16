package main

import "testing"

func TestGreedy(t *testing.T) {
	type args struct {
		list []int
		X    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "500円+100円=600円のケース",
			args: args{
				list: []int{2, 1, 1, 1, 1, 1},
				X:    600,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greedy(tt.args.list, tt.args.X); got != tt.want {
				t.Errorf("Greedy() = %v, want %v", got, tt.want)
			}
		})
	}
}
