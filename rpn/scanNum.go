package rpn

//import "fmt"

func scanNum(statement []rune, i *int, rpn []string) []string {
	iStart := *i
	for ; *i < len(statement) && (statement[*i] >= '0' && statement[*i] <= '9'); *i++ {
	}
	rpn = append(rpn, string(statement[iStart:*i]))
	*i--
	//fmt.Print(string(statement[iStart:i]))
	return rpn
}
