package main

import "testing"

func TestFindMinimumPair(t *testing.T) {
	type args struct {
		list []int
		K    int
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
				list: []int{1, 2, 3, 4},
				K:    4,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMinimumPair(tt.args.list, tt.args.K); got != tt.want {
				t.Errorf("FindMinimumPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
