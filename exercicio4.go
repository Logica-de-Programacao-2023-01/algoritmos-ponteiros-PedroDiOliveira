package main

import "strconv"

func adicionedoisdigitos(ptr *int) {
	str := strconv.Itoa(*ptr)
	digits := make([]int, len(str))

	for i, char := range str {
		digit, _ := strconv.Atoi(string(char))
		digits[i] = digit
	}
	x := len(digits)
	*ptr = digits[x-1] + digits[x-2]
}
