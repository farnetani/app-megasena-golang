package utils

import (
	"reflect"

	"github.com/farnetani/app-megasena-golang/models"
)

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

func eq(a, b models.Combinations) bool {
	return reflect.DeepEqual(a.Numbers, b.Numbers)
}

func Contains(list []models.Combinations, x models.Combinations) bool {
	for i := range list {
		if eq(x, list[i]) {
			return true
		}
	}
	return false
}
