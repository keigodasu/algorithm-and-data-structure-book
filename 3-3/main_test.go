package main

import "testing"

func TestFindMin(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name    string
		args    args
		wantNum int
	}{
		// TODO: Add test cases.
		{name: "found min val", args: args{list: []int{1, 2, 3}}, wantNum: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNum := FindMin(tt.args.list); gotNum != tt.wantNum {
				t.Errorf("FindMin() = %v, want %v", gotNum, tt.wantNum)
			}
		})
	}
}
