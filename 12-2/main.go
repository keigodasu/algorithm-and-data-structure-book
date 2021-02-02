package main

func MergeSort(list []int) []int  {
	var mergedList []int
	if len(list) < 2{
		return list
	}

	mid := int(len(list) / 2)

	ml := MergeSort(list[mid:])
	mr := MergeSort(list[:mid])

	i, j := 0, 0
	for i < len(mr) && j < len(ml) {
		if mr[i] > ml[j] {
			mergedList = append(mergedList, ml[j])
			j++
		} else {
			mergedList = append(mergedList, mr[i])
			i++
		}
	}

	mergedList = append(mergedList, mr[i:]...)
	mergedList = append(mergedList, ml[j:]...)

	return mergedList
}
