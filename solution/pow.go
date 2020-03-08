package solution

func pow(base float64, pow float64) float64 {
	if pow == 0 {
		return 1
	}
	powInt := int(pow)
	k := base
	for ; powInt > 1; powInt-- {
		base *= k
	}
	return base
}
