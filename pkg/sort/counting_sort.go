package sort

func CountingSort(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}
	max := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
		}
	}

	count := make([]int, max+1)
	for _, value := range slice {
		count[value]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	result := make([]int, len(slice))
	for i := len(slice) - 1; i >= 0; i-- {
		result[count[slice[i]]-1] = slice[i]
		count[slice[i]]--
	}
	return result
}
