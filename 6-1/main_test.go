package main

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		key  int
		list []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"found case", args{key: 4, list: []int{1, 3, 4, 5, 8, 10}}, 2},
		{"not found case", args{key: 12, list: []int{1, 3, 4, 5, 8, 10}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.key, tt.args.list); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
