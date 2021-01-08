package main

import "testing"

func TestPartialSum(t *testing.T) {
	type args struct {
		i    int
		w    int
		list []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "found",
			args: args{
				i:    5,
				w:    5,
				list: []int{1, 2, 3, 4, 6},
			},
			want: true,
		},
		{
			name: "not found",
			args: args{
				i:    2,
				w:    5,
				list: []int{1, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartialSum(tt.args.i, tt.args.w, tt.args.list); got != tt.want {
				t.Errorf("PartialSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
