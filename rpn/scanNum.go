package rpn

//import "fmt"

func scanNum(statement []rune, i *int, rpn []string) []string {
	iStart := *i
	if statement[*i] == '-' {
		*i++
	}
	for ; *i < len(statement) && (IsNum(statement[*i]) || IsSep(statement[*i])); *i++ {
	}
	//fmt.Print(string(statement[iStart:*i]))
	//fmt.Println()
	//fmt.Println(rpn)
	rpn = append(rpn, string(statement[iStart:*i]))
	//fmt.Println(rpn)
	*i--
	return rpn
}
