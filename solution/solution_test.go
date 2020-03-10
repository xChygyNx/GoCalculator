package solution

import "testing"
import "errors"

type powTest struct {
	base float64
	pow  float64
	res  float64
	err  error
}

type resTest struct {
	rpn []string
	res float64
	err error
}

var zinp = errors.New("Zero in negative pow")
var dbz = errors.New("Division by 0")
var is = errors.New("Incorrect statement")

var powTestsCorrect = []powTest{
	{2, 2, 4, nil},
	{4, -1, 0.25, nil},
	{1.5, 2, 2.25, nil},
	{0, 0, 1, nil},
	{0, 20, 0, nil},
	{4, 6, 4096, nil},
}

var powTestsErr = []powTest{
	{0, -3, 0, zinp},
	{0, -100, 0, zinp},
}

var resTestsCorrect = []resTest{
	{[]string{"42", "26", "+"}, 68, nil},
	{[]string{"2", "2", "2", "*", "+"}, 6, nil},
	{[]string{"100", "33", "-", "0", "0", "^", "-"}, 66, nil},
	{[]string{"3", "17", "*", "3", "-", "2", "9", "5", "-", "^", "/"}, 3, nil},
}

var resTestsErr = []resTest{
	{[]string{"10", "5", "5", "-", "/"}, 0, dbz},
	{[]string{"825", "16", "-35", "*"}, 0, is},
	{[]string{"12", "-3", "/", "^"}, 0, is},
}

func TestPow(t *testing.T) {
	for _, elem := range powTestsCorrect {
		got, err := pow(elem.base, elem.pow)
		if got != elem.res || elem.err != err {
			t.Error(
				"For", elem.base, "^", elem.pow,
				"Want", elem.res, "error", elem.err,
				"Got", got, "error", err,
			)
		}
	}
	for _, elem := range powTestsErr {
		got, err := pow(elem.base, elem.pow)
		if got != elem.res || err.Error() != elem.err.Error() {
			t.Error(
				"For", elem.base, "^", elem.pow,
				"Want", elem.res, "error", elem.err,
				"Got", got, "error", err,
			)
		}
	}
}

func TestRes(t *testing.T) {
	for _, elem := range resTestsCorrect {
		got, err := Res(elem.rpn)
		if got != elem.res || elem.err != err {
			t.Error(
				"For", elem.rpn,
				"Want", elem.res, "Error", elem.err,
				"Got", got, "Error", err,
			)
		}
	}
	for _, elem := range resTestsErr {
		got, err := Res(elem.rpn)
		if got != elem.res || elem.err.Error() != err.Error() {
			t.Error(
				"For", elem.rpn,
				"Want", elem.res, "Error", elem.err,
				"Got", got, "Error", err,
			)
		}
	}
}
