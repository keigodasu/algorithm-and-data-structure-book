package main

import "testing"

func TestSectionSchedule(t *testing.T) {
	type args struct {
		tasks []Task
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case01",
			args: args{tasks: []Task{
				{
					start: 1,
					end:   2,
				},
				{
					start: 2,
					end:   6,
				},
				{
					start: 5,
					end:   7,
				},
			}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SectionSchedule(tt.args.tasks); got != tt.want {
				t.Errorf("SectionSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}
