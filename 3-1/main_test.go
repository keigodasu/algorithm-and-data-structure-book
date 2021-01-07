package main

import "testing"

func TestFindNumber(t *testing.T) {
	type args struct {
		list []int
		v    int
	}
	tests := []struct {
		name          string
		args          args
		wantExistFlag bool
	}{
		{
			name: "found",
			args: args{
				list: []int{1,2,3},
				v:    1,
			},
			wantExistFlag: true,
		},
		{
			name: "not found",
			args: args{
				list: []int{1,2,3},
				v:    0,
			},
			wantExistFlag: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotExistFlag := findNumber(tt.args.list, tt.args.v); gotExistFlag != tt.wantExistFlag {
				t.Errorf("findNumber() = %v, want %v", gotExistFlag, tt.wantExistFlag)
			}
		})
	}
}
