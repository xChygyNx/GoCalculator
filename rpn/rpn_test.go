package rpn

import "testing"

type testTransform struct {
	statement string
	rpn       []string
}

type testCloseParenthese struct {
	rpnStart  []string
	opsStart  []rune
	rpnFinish []string
	opsFinish []rune
}

var testsTransform = []testTransform{
	{"2 + 2 * 2", []string{"2", "2", "2", "*", "+"}},
	{"(15-6) / 2", []string{"15", "6", "-", "2", "/"}},
	{"0 / 0", []string{"0", "0", "/"}},
	{" 8 - - 3", []string{"8", "-", "3", "-"}},
}

var testsCloseParenthese = []testCloseParenthese{
	{[]string{"-5", "8", "+", "6", "4"}, []rune{'(', '-', '*'}, []string{"-5", "8", "+", "6", "4", "*", "-"}, []rune{}},
	{[]string{"2", "152", "3", "4", "6", "^", "2"}, []rune{'+', '*', '(', '+', '/'}, []string{"2", "152", "3", "4", "6", "^", "2", "/", "+"}, []rune{'+', '*'}},
}

func TestTransform(t *testing.T) {
	for _, elem := range testsTransform {
		if got := Transform(elem.statement); !eqSliceStr(got, elem.rpn) {
			t.Error(
				"For", elem.statement,
				"Want", elem.rpn,
				"Got", got,
			)
		}
	}
}

func TestcloseParenthese(t *testing.T) {
	for _, elem := range testsCloseParenthese {
		gotRpn, gotOps := closeParenthese(elem.rpnStart, elem.opsStart)
		if !eqSliceStr(gotRpn, elem.rpnFinish) || !eqSliceRune(gotOps, elem.opsFinish) {
			t.Error(
				"For", elem.rpnStart, "&", elem.opsStart,
				"Want", elem.rpnFinish, "&", elem.opsFinish,
				"Got", gotRpn, "&", gotOps,
			)
		}
	}
}

func eqSliceStr(s1 []string, s2 []string) bool {
	var i int
	for ; i < len(s1); i++ {
		if i == len(s2) || s1[i] != s2[i] {
			return false
		}
	}
	if i != len(s1) {
		return false
	}
	return true
}

func eqSliceRune(s1 []rune, s2 []rune) bool {
	var i int
	for ; i < len(s1); i++ {
		if i == len(s2) || s1[i] != s2[i] {
			return false
		}
	}
	if i != len(s1) {
		return false
	}
	return true
}
