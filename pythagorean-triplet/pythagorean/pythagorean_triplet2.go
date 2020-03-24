package pythagorean

import (
	"math"
)

type Triplet [3]int

func Sum2(p int) []Triplet {
	var a, b int = 1, 2
	res := make([]Triplet, 0)
	for ; a < b; a++ {
		for b = 1; math.Pow(float64(a), 2)+math.Pow(float64(b), 2) < math.Pow(float64(p-a-b), 2); b++ {

		}
		if math.Pow(float64(a), 2)+math.Pow(float64(b), 2) == math.Pow(float64(p-a-b), 2) {
			tr := sortTriplet(a, b, p)
			res = append(res, tr)
		}
	}
	return res
}

func Range2(min, max int) []Triplet {
	var res, temp []Triplet
	down := bottomBorder(min)
	up := upperBorder(max)
	for i := down; i <= up; i++ {
		temp = Sum2(i)
		if checkBorder(temp, min, max) {
			res = append(res, temp...)
		}
	}
	return res
}

func bottomBorder(min int) int {
	gipotenuza := int(math.Sqrt(2 * math.Pow(float64(min), 2)))
	return gipotenuza + 2*min
}

func upperBorder(max int) int {
	katet := int(math.Sqrt(math.Pow(float64(max), 2) / 2))
	return max + 2*katet
}

func sortTriplet(a, b, p int) Triplet {
	var res Triplet
	res[0] = a
	res[1] = b
	res[2] = p - a - b
	return res
}

func checkBorder(sets []Triplet, min int, max int) bool {
	for _, tr := range sets {
		for i := 0; i < 3; i++ {
			if tr[i] < min || tr[i] > max {
				return false
			}
		}
	}
	return true
}
