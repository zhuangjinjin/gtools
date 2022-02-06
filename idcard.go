package gtools;

import (
	"strconv"
)

var factors = []int {7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var masks = []int {1, 0, 10, 9, 8, 7, 6, 5, 4, 3, 2}

func Verificate(idCardStr string) bool {
	if (len(idCardStr) != 18) {
		return false
	}
	var sum int
	for i := 0; i < 17; i++ {
		digit, _ := strconv.Atoi(idCardStr[i:i+1])
		sum += digit * factors[i]
	}
	calMaskIndex := sum % 11
	var mask int
	if (idCardStr[17:] == "X") {
		mask = 10
	} else {
		mask, _ = strconv.Atoi(idCardStr[17:])
	}

	if mask == masks[calMaskIndex] {
		return true
	} else {
		return false
	}
}