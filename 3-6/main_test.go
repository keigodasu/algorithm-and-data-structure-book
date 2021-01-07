package main

import "testing"

func TestFindSum(t *testing.T) {
	type args struct {
		list []int
		num  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "found",
			args: args{
				list: []int{1, 2, 3, 4, 5},
				num:  9,
			},
			want: true,
		},
		{
			name: "not found",
			args: args{
				list: []int{1, 2, 3, 4},
				num:  19,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSum(tt.args.list, tt.args.num); got != tt.want {
				t.Errorf("FindSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
