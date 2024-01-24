package stringmatching

//I don't know what it's called, so i call it "unknown"
func UnknownMatching(str, substr string) (index int, found bool) {
	for i, j := 0, 0; i < len(str); i++ {
		if str[i] == substr[j] {
			j++
			if j == len(substr) {
				return i + 1 - j, true
			}
		} else {
			j = 0
		}
	}
	return -1, false
}
