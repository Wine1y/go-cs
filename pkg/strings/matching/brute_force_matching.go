package stringmatching

func BruteForceMatching(str, substr string) (index int, found bool) {
	for i := 0; i <= len(str)-len(substr); i++ {
		offset := 0
		for offset < len(substr) && str[i+offset] == substr[offset] {
			offset++
		}
		if offset == len(substr) {
			return i, true
		}
	}
	return -1, false
}
