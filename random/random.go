package random

import (
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RangeNumbers: result a random range of numbers no repeated as a slice within a range
func RangeNumbers(min int, max int, n int) []int {
	arr := make([]int, n)
	var r, i int
	for r = 0; r <= n-1; r++ {
		nr := rand.Intn(max) + min
		exists := false
		for i = 0; i < len(arr); i++ {
			if arr[i] == nr {
				exists = true
				break
			}
		}
		if !exists {
			arr[r] = nr
		} else {
			r--
		}
	}
	sort.Ints(arr)
	return arr
}
