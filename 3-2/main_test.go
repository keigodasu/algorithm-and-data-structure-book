package main

import "testing"

func TestFindIndex(t *testing.T) {
	type args struct {
		list   []int
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"found",
			args{
				list:   []int{1, 2, 3},
				number: 1,
			},
			0,
		},
		{
			"not found",
			args{
				list:   []int{1, 2, 3},
				number: 0,
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindIndex(tt.args.list, tt.args.number); got != tt.want {
				t.Errorf("FindIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
