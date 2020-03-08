package format

import (
	"strconv"
)

func AccuracyOut(num float64) int {
	numInt := int(num)
	numIntStr := strconv.Itoa(numInt)
	numStr := strconv.FormatFloat(num, 'f', -1, 64)
	if len(numStr) == len(numIntStr) {
		return 0
	}
	return len(numStr) - len(numIntStr) - 1
}
