package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/farnetani/app-megasena-golang/megasena"
)

func main() {
	qt := GetParams()

	// Return models.Combinations and Print in console and save in numbers.txt file
	megasena.ListAndSaveCombinations(qt, 6)
}

// GetParams : return the quantity of games in megasena
func GetParams() int {
	if len(os.Args) > 1 {
		paramsNumbers := os.Args[1]
		qt, err := strconv.Atoi(paramsNumbers)
		if err != nil {
			panic(fmt.Errorf("error in parameters %s", err.Error()))
		}
		return qt
	}
	return 1 // return one combination if not given in parameters
}
