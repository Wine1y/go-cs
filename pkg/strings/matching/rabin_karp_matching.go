package stringmatching

import (
	"hash/adler32"
)

var hashFunction func([]byte) uint32 = adler32.Checksum

func RabinKarpMatching(str, substr string) (index int, found bool) {
	subStrHash := hashFunction([]byte(substr))
	for i := 0; i <= len(str)-len(substr); i++ {
		hash := hashFunction([]byte(str[i : i+len(substr)]))
		if hash == subStrHash && str[i:i+len(substr)] == substr {
			return i, true
		}
	}
	return -1, false
}
