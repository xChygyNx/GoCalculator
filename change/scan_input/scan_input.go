package scan_input

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

func ScanCoins(nominals string) ([]int, error) {
	coinsTabStr := strings.Split(nominals, " ")
	coinsTab := make([]int, 0, len(coinsTabStr))
	for _, elem := range coinsTabStr {
		num, err := strconv.Atoi(elem)
		if err != nil || num < 1 {
			return []int{}, errors.New("Invalid nominal of coin, need integer more 0")
		}
		coinsTab = append(coinsTab, num)
	}
	sort.Ints(coinsTab)
	return coinsTab, nil
}

func ScanTarget(str string) (int, error) {
	target, err := strconv.Atoi(str)
	if err != nil || target < 0 {
		return 0, errors.New("Invalid target, need positive integer or 0")
	}
	return target, nil
}
