package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"

	"github.com/alok87/goutils/pkg/random"
)

type combinations struct {
	numbers []int
}

func main() {
	// 10:18h

	q := 200
	numbers := make([]combinations, q)

	for i := 1; i <= q; i++ {
		r := random.RangeInt(1, 60, 6)
		arrNumbersUnique := Uniques(r) // return only unique
		// Valid if numbers is not repeated
		if len(arrNumbersUnique) == 6 {
			sort.Ints(arrNumbersUnique)
			numbers[i-1] = combinations{arrNumbersUnique}
		} else {
			i = i - 1 // decrement because numbers repeateds
		}
	}

	var uniqueNumbers []combinations
	// add the last occurrence always
	for i := len(numbers) - 1; i >= 0; i-- {
		if !contains(uniqueNumbers, numbers[i]) {
			// "append" to the front of the array
			uniqueNumbers = append([]combinations{numbers[i]}, uniqueNumbers...)
		}
	}

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
		for g, n := range combinations.numbers {
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
}

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Uniques(arr []int) []int {

	counter := map[int]int{}
	for _, item := range arr {
		counter[item]++
	}

	resp := make([]int, 0, len(counter))
	for key, value := range counter {
		if value == 1 {
			item := key
			resp = append(resp, item)
		}
	}

	return resp
}

func RemoveDuplicateValues(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func UniqueSlice(s []int) []int {
	inResult := make(map[int]bool)
	var result []int
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

func eq(a, b combinations) bool {
	return reflect.DeepEqual(a.numbers, b.numbers)
}

func contains(list []combinations, x combinations) bool {
	for i := range list {
		if eq(x, list[i]) {
			return true
		}
	}
	return false
}
