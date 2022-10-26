// Package SCC or Shortcut codes is a small library of handwritten code to provide simple code
// shortcuts that can be used for a wide range of functions or programs.
package SCC

import (
	"errors"
	"math"
)

// NumBreaker takes an integer and breaks it into it's places (hundreds|tens|ones) it only takes numBreak up to the millions.
// It returns an error if and only if the inputted number becomes or is negative.
func NumBreaker(num int) ([]int, error) {
	var numBreak []int
	var m, cth, xth, th, c, t int
	for {
		switch {
		case num >= 1000000:
			m = num / 1000000
			numBreak = append(numBreak, m*1000000)
			num %= 1000000
		case num >= 100000:
			cth = num / 100000
			numBreak = append(numBreak, cth*100000)
			num %= 100000
		case num >= 10000:
			xth = num / 10000
			numBreak = append(numBreak, xth*10000)
			num %= 10000
		case num >= 1000:
			th = num / 1000
			numBreak = append(numBreak, th*1000)
			num %= 1000
		case num >= 100:
			c = num / 100
			numBreak = append(numBreak, c*100)
			num %= 100
		case num >= 10:
			t = num / 10
			numBreak = append(numBreak, t*10)
			num %= 10
		case num >= 1:
			numBreak = append(numBreak, num)
			num = 0
		case num < 0:
			return []int{}, errors.New("num is negative, check switch")
		default:
			return numBreak, nil
		}
	}
}

// NumBreaker takes an integer and breaks it into it's places (hundreds|tens|ones) it takes up to the max value of an int.
// It returns an error if and only if the inputted number becomes or is negative.
func NumBreaker2(num int) ([]int, error) {
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

func NumBreaker3(num int) ([]int, error) {
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
