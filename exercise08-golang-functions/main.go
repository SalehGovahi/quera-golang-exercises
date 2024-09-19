package main

import (
	"math"
	"strconv"
)

type FilterFunc func(int) bool
type MapperFunc func(int) int

func IsSquare(x int) bool {
	if x < 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(x)))
	return sqrt*sqrt == x
}

func IsPalindrome(x int) bool {
	if x < 0 {
		x = -x
	} 
	strNum := strconv.Itoa(x)
	length := len(strNum)
	for i := 0; i < length/2; i++ {
		if strNum[i] != strNum[length-i-1] {
			return false
		}
	}
	return true
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Cube(num int) int {
	return num * num * num
}

func Filter(input []int, f FilterFunc) []int {
	var result []int
	for _, number := range input {
		if f(number) {
			result = append(result, number)
		}
	}
	return result
}

func Map(input []int, m MapperFunc) []int {
	for index, number := range input {
		input[index] = m(number)
    }
	return input
}