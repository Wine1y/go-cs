package search

func InterpolationSearch(slice []int, item int) (index int, found bool) {
	for low, high := 0, len(slice)-1; low <= high; {
		mid := low + (((item - slice[low]) * (high - low)) / (slice[high] - slice[low]))
		if slice[mid] == item {
			return mid, true
		}
		if slice[mid] > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1, false
}
