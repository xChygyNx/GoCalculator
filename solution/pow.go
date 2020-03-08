package solution

func pow(base float64, pow float64) float64 {
	powInt := int(pow)
	if powInt == 0 {
		return 1
	}
	k := base
	if powInt > 0 {
		for ; powInt > 1; powInt-- {
			base *= k
		}
	} else {
		for ; powInt <= 0; powInt++ {
			base /= k
		}
	}
	return base
}
