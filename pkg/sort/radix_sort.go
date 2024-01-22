package sort

func countingSortExp(slice []int, exp int) []int {
	count := make([]int, 10)
	for i := 0; i < len(slice); i++ {
		digit := slice[i] / exp % 10
		count[digit]++
	}

	for i := 0; i <= 10; i++ {
		count[i] += count[i-1]
	}

	result := make([]int, len(slice))
	for i := len(slice) - 1; i >= 0; i-- {
		digit := slice[i] / exp % 10
		result[count[digit]-1] = slice[i]
		count[digit]--
	}
	return result
}

func RadixSort(slice []int) {
	if len(slice) == 0 {
		return
	}
	max := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
		}
	}

	for exp := 1; max/exp >= 1; exp *= 10 {
		slice = countingSortExp(slice, exp)
	}
}
