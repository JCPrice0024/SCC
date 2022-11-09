// Package SCC or Shortcut codes is a small library of handwritten code to provide simple code
// shortcuts that can be used for a wide range of functions or programs.
package SCC

import (
	"errors"
	"log"
	"math"
	"strings"
)

const (
	Thousand     = 1e3
	Million      = 1e6
	Billion      = 1e9
	Trillion     = 1e12
	Quadrillion  = 1e15
	Quintilliion = 1e18
)

// NumBreaker takes an integer and breaks it into it's places (hundreds|tens|ones) it takes up to the max value of an int.
// It returns an error if and only if the inputted number becomes or is negative.
func NumBreaker(num int) ([]int, error) {
	var numBreak []int
	if num < 0 {
		return []int{}, errors.New("num was negative, check input")
	}
	for i := 18; i >= 0; i-- {
		exp := int(math.Pow10(i))
		if num/exp < 1 {
			continue
		}
		numBreak = append(numBreak, (num/exp)*exp)
		num %= exp
	}
	if num < 0 {
		return []int{}, errors.New("num became negative, check loop")
	}
	return numBreak, nil
}

// NumBreaker2 takes an integer and breaks it into it's places (hundreds|tens|ones) it takes up to the max value of an int.
// It returns the numbers in 1|2|3 form as opposed to 100|20|3
// It returns an error if and only if the inputted number becomes or is negative.
func NumBreaker2(num int) ([]int, error) {
	var numBreak []int
	if num < 0 {
		return []int{}, errors.New("num was negative, check input")
	}
	for num > 0 {
		remainder := num % 10
		numBreak = append(numBreak, remainder)
		num /= 10
	}
	for i, j := 0, len(numBreak)-1; i < j; i, j = i+1, j-1 {
		numBreak[i], numBreak[j] = numBreak[j], numBreak[i]
	}
	return numBreak, nil
}

// NumGrouper takes a slice of broken up ints and groups them by their number of zeros. It groups accurately up to 10^18.
// For example all the numbers ending in Thousand and less would look like: [100000, 10000, 1000],[100, 10, 1])
func NumGrouper(nums []int) [][]int {
	var groups [][]int
	var qt, qd, tr, b, m, th, h []int
	for _, v := range nums {
		switch {
		case v >= Quintilliion:
			qt = append(qt, v)
		case v < Quintilliion && v >= Quadrillion:
			qd = append(qd, v)
		case v < Quadrillion && v >= Trillion:
			tr = append(tr, v)
		case v < Trillion && v >= Billion:
			b = append(b, v)
		case v < Billion && v >= Million:
			m = append(m, v)
		case v < Million && v >= Thousand:
			th = append(th, v)
		case v < Thousand:
			h = append(h, v)
		}
	}
	groups = append(groups, qt, qd, tr, b, m, th, h)
	return groups
}

// NumSplitter is an extension for NumGrouper. It passes NumBreaker an int, and NumGrouper the
// resultingslice of int. It then ranges over the result of NumGrouper and splits them by their number of zeros
// into one readable slice of int for NumReader.
func NumSplitter(num int) []int {
	split := []int{}
	n, err := NumBreaker(num)
	if err != nil {
		log.Println(err)
		return []int{}
	}
	nums := NumGrouper(n)
	for i := range nums {
		split = rangeSplitter(nums[i], split)
	}
	return split
}

// rangeSplitter is an extension of NumSplitter. It takes in one of the slices of int ranged over by NumSplitter (nums)
// and ranges over it while perfroming operations on their values to make a readable slice for NumReader.
// (split is the output slice made in NumSplitter)
func rangeSplitter(nums, split []int) []int {
	n := 18
	exp := int(math.Pow10(n))
	for i, v := range nums {
		if v == 0 && i == len(nums)-1 && exp != 1 {
			split = append(split, exp)
			continue
		} else if v == 0 {
			return split
		}
		exp = validExp(v, exp, n)
		switch {
		case i < len(nums)-1 && v/exp == 10:
			split = append(split, (nums[i]/exp)+(nums[i+1]/exp))
			nums[i+1] = 0
		case v/exp >= 1:
			split = append(split, v/exp)
		}
		if i == len(nums)-1 && exp != 1 {
			split = append(split, exp)
		}
	}
	return split
}

// validExp takes in a power of 10 exponent (exp) raised to a certain power (n) and tests
// that the exponent has the correct number of zeroes to perform an operation on the given value
// if it doesn't it reduces the zeroes by three
func validExp(val, exp, n int) int {
	for val/exp < 1 {
		n -= 3
		exp = int(math.Pow10(n))
	}
	return exp
}

// NumReader takes in an int value and returns that int in english words.
// NumReader does not take negative values or values greater than the max value of an int.
func NumReader(num int) string {
	if num == 0 {
		return "Zero"
	}
	n := NumSplitter(num)
	m := map[int]string{
		Quintilliion: "Quintilliion", Quadrillion: "Quadrillion", Trillion: "Trillion", Billion: "Billion", Million: "Million", Thousand: "Thousand", 100: "Hundred", 10: "Ten",
		1: "One", 2: "Two", 3: "Three", 4: "Four", 5: "Five", 6: "Six", 7: "Seven", 8: "Eight", 9: "Nine", 11: "Eleven", 12: "Twelve",
		13: "Thirteen", 14: "Fourteen", 15: "Fifteen", 16: "Sixteen", 17: "Seventeen", 18: "Eighteen", 19: "Nineteen",
		20: "Twenty", 30: "Thirty", 40: "Forty", 50: "Fifty", 60: "Sixty", 70: "Seventy", 80: "Eighty", 90: "Ninety",
	}
	s := createNumString(n, m)
	return s
}

// createNumString is an extension of NumReader. It simply makes a string using the map values from
// NumReader and the values stored in our slice of int. It prints the numbers as they come so the slice
// looks like this: [100 20 3 1000 100 20 3]
func createNumString(n []int, m map[int]string) string {
	var str strings.Builder
	for i, v := range n {
		switch {
		case v >= 100 && v < 1000 && i == len(n)-1:
			str.WriteString(m[v/100] + " " + m[100])
			return str.String()
		case v >= 100 && v < 1000:
			str.WriteString(m[v/100] + " " + m[100] + " ")
		case i == len(n)-1:
			str.WriteString(m[v])
			return str.String()
		default:
			str.WriteString(m[v] + " ")
		}
	}
	return str.String()
}

// BinarySearch takes an int and an array of sorted ints and finds x by comparing the midpoints of the array.
// It goes left if x is less than our midpoint and right if it's greater. It has O(log(n)) time complexity.
func BinarySearch(x int, nums []int) int {
	low := 0
	high := len(nums)
	mid := 0
	for low != high {
		mid = (low + high) / 2
		if nums[mid] == x {
			return mid
		}
		if nums[mid] > x {
			high = mid - 1
			continue
		}
		if nums[mid] < x {
			low = mid + 1
			continue
		}
	}
	if low == high {
		return high
	}
	if mid == len(nums)-1 && nums[mid] != x {
		return -1
	}
	return mid
}
