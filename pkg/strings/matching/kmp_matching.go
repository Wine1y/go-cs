package stringmatching

func getPrefixTable(str string) []int {
	table := make([]int, len(str))
	for i, j := 1, 0; i < len(str); {
		if str[i] == str[j] {
			table[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = table[j-1]
		} else {
			table[i] = 0
			i++
		}
	}
	return table
}

func KnuthMorrisPrattMatching(str, substr string) (index int, found bool) {
	prefixTable := getPrefixTable(substr)
	for i, j := 0, 0; i < len(str); i++ {
		if str[i] == substr[j] {
			if j == len(substr)-1 {
				return i - j, true
			} else {
				j++
			}
		} else if j > 0 {
			j = prefixTable[j-1]
		}
	}
	return -1, false
}
