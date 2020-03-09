package solution

import "testing"
import "errors"

type powTest struct {
	base float64
	pow  float64
	res  float64
	err  error
}

var zinp = errors.New("Zero in negative pow")

var powTests = []powTest{
	{2, 2, 4, nil},
	{4, -1, 0.25, nil},
	{1.5, 2, 2.25, nil},
	{0, 0, 1, nil},
	{0, 20, 0, nil},
	{4, 6, 4096, nil},
	{0, -3, 0, zinp},
}

func TestPow(t *testing.T) {
	for _, elem := range powTests {
		got, err := pow(elem.base, elem.pow)
		if got != elem.res || errors.As(err, zinp) {
			t.Error(
				"For", elem.base, "^", elem.pow,
				"Want", elem.res, "error", elem.err,
				"Got", got, "error", err,
			)
		}
	}
}
