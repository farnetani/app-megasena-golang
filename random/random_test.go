package random

import (
	"fmt"
	"sort"
	"testing"

	"github.com/farnetani/app-megasena-golang/models"
	"github.com/farnetani/app-megasena-golang/utils"
	"github.com/stretchr/testify/assert"
)

func TestRandom(t *testing.T) {
	tests := []struct {
		name     string
		property any
		expected string
	}{
		{
			name:     "One test...",
			property: "1",
			expected: "2",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			n := RangeNumbers(1, 60, 6)
			fmt.Println("Numbers", n)
			assert.NotNil(t, n, n)
		})
	}
}

func TestRangeNumbersRepeated(t *testing.T) {
	q := 1000
	numbers := make([]models.Combinations, q)

	var arrNumbersUnique []int
	var i, c int
	for i = 0; i < q; i++ {
		r := RangeNumbers(1, 60, 6)
		fmt.Println(r)
		arrNumbersUnique = utils.Uniques(r)
		if len(arrNumbersUnique) == 6 {
			sort.Ints(arrNumbersUnique)
			numbers[i] = models.Combinations{Numbers: arrNumbersUnique}
			c++
		}
	}
	assert.Equal(t, q, c)
}

func TestFailtUniqueValues(t *testing.T) {
	numbers := GenerateNumbers(1, 60, 6, 1000)
	for _, n := range numbers {
		arrUnique := utils.Uniques(n.Numbers)
		if len(arrUnique) < 6 {
			assert.Fail(t, "Values repeated")
		}
	}
}

func TestFailUniqueGroupNumbers(t *testing.T) {
	numbers := GenerateNumbers(1, 60, 6, 1000)
	var uniqueNumbers []models.Combinations
	for i := len(numbers) - 1; i >= 0; i-- {
		if !utils.Contains(uniqueNumbers, numbers[i]) {
			uniqueNumbers = append([]models.Combinations{numbers[i]}, uniqueNumbers...)
		} else {
			assert.Fail(t, "Group of numbers repeated")
		}
	}
}
