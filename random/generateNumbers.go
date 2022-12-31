package random

import (
	"fmt"
	"log"
	"os"

	"github.com/farnetani/app-megasena-golang/models"
	"github.com/farnetani/app-megasena-golang/utils"
)

func GenerateNumbers(min int, max int, sequences int, q int) []models.Combinations {
	numbers := make([]models.Combinations, q)

	for i := 1; i <= q; i++ {
		arrNumbersUnique := RangeNumbers(min, max, sequences)
		numbers[i-1] = models.Combinations{Numbers: arrNumbersUnique}
	}

	var uniqueNumbers []models.Combinations
	// add the last occurrence always
	for i := len(numbers) - 1; i >= 0; i-- {
		if !utils.Contains(uniqueNumbers, numbers[i]) {
			// "append" to the front of the array
			uniqueNumbers = append([]models.Combinations{numbers[i]}, uniqueNumbers...)
		}
	}

	// create the file numbers.txt
	f, err := os.Create("numbers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Print combination
	f.WriteString("SYSTEM SUPER-MEGA-SENA-DA-VIRADA")
	f.WriteString("\n----------------------------------------\n")
	f.WriteString("Copyright (c) Arlei F. Farnetani Junior")
	f.WriteString("\n----------------------------------------\n")
	f.WriteString(fmt.Sprintf("%d Combinações para Mega-Sena da Virada", q))
	f.WriteString("\n----------------------------------------\n")
	for i, combinations := range uniqueNumbers {
		s := ""
		fmt.Printf("#%d | ", i+1)
		s = fmt.Sprintf("#%d | ", i+1)
		for g, n := range combinations.Numbers {
			fmt.Printf(" %d ", n)
			s = s + fmt.Sprintf(" %d ", n)
			if g < 5 {
				fmt.Print("-")
				s = s + "-"
			}
		}
		fmt.Println("\n----------------------------------------")
		s = s + "\n----------------------------------------\n"
		_, err = f.WriteString(s)
		if err != nil {
			log.Fatal(err)
		}
	}
	return numbers
}
