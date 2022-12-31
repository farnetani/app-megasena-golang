package megasena

import (
	"github.com/farnetani/app-megasena-golang/models"
	"github.com/farnetani/app-megasena-golang/random"
)

func ListAndSaveCombinations(q int, n int) []models.Combinations {
	return random.GenerateNumbers(1, 60, n, q)
}
