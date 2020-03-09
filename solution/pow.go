package solution

// Функция возведения в степень

import "errors"

func pow(base float64, pow float64) (float64, error) {
	powInt := int(pow)
	if powInt == 0 {
		return 1, nil
	}
	k := base
	if powInt > 0 {
		for ; powInt > 1; powInt-- {
			base *= k
		}
	} else if base != 0 {
		for ; powInt <= 0; powInt++ {
			base /= k
		}
	} else {
		return 0, errors.New("Zero in negative pow")
	}
	return base, nil
}
