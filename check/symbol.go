package check

import . "../rpn"

func checkSymbol(r rune) bool {
	if IsNum(r) || IsOp(r) || IsParenthes(r) || IsSpace(r) || IsSep(r) {
		return true
	}
	return false
}
