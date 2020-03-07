package solution

func pow(base int, pow int) int {
	if pow == 0 {
		return 1
	}
	k := base
	for ; pow > 1; pow-- {
		base *= k
	}
	return base
}
