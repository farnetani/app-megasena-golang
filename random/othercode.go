package random

import (
	"math/rand"
	"time"
)

func GenerateRandomNumbers(q int, max int) []int {
	rand.Seed(time.Now().UnixNano())
	var numbers []int
	for len(numbers) < q {
		n := rand.Intn(max) + 1
		exists := false
		for _, num := range numbers {
			if num == n {
				exists = true
				break
			}
		}
		if !exists {
			numbers = append(numbers, n)
		}
	}
	return numbers
}
