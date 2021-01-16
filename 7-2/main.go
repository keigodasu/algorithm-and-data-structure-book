package main

import "sort"

type Task struct {
	start, end int
}

func SectionSchedule(list []Task) int {
	sort.Slice(list, func(i, j int) bool {
		return list[i].end < list[j].end
	})

	res, curentEndTime := 0, 0

	for i := 0; i < len(list); i++ {
		if list[i].start < curentEndTime {
			continue
		}
		res++
		curentEndTime = list[i].end
	}

	return res
}
