package main

import (
	"reflect"
	"testing"
)

func TestMergeSort1(t *testing.T) {
	type args struct {
		list  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test",
			args: args{
				list:  []int{3,4,1,9},
			},
			want: []int{1,3,4,9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}