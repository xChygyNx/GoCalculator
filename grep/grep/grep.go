package grep

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var flags = [5]byte{'x', 'n', 'i', 'v', 'l'}

func Grep() {
	if len(os.Args) < 3 {
		Usage()
		return
	}
	i := 1
	flags, err := parseFlags(os.Args[1], &i)
	if err != nil {
		Usage()
		fmt.Printf("\n%v\n", err)
		return
	}
	pattern := os.Args[i]
	i++
	if i == len(os.Args) {
		Usage()
		return
	}
	files := make([]string, 0)
	for ; i < len(os.Args); i++ {
		files = append(files, os.Args[i])
	}
	res := Search(pattern, flags, files)
	for _, elem := range res {
		fmt.Println(elem)
	}
}

func Search(pattern string, flags []string, files []string) []string {
	res := make([]string, 0)
	flagsMap := flagsMap(flags)
	fileContent := readFiles(files)
	if flagsMap["-l"] {
		res = fileNames(flagsMap, pattern, fileContent)
	} else {
		res = strMatch(flagsMap, pattern, fileContent, len(files) > 1)
	}
	return res
}

// create map of flags fo more convenient subsequent work
func flagsMap(flags []string) map[string]bool {
	res := make(map[string]bool)
	for _, flag := range flags {
		res[flag] = true
	}
	return res
}

// Create map of file content flom list of
// names of files
func readFiles(files []string) map[string][]string {
	res := make(map[string][]string)
	for _, elem := range files {
		file, err := os.Open(elem)
		if err == nil {
			bufRead := bufio.NewReader(file)
			value := make([]string, 0)
			var readError bool
			for str, errRead := bufRead.ReadString('\n'); ; str, errRead = bufRead.ReadString('\n') {
				if errRead == nil {
					if str != "" && str != "\n" {
						value = append(value, strings.Trim(str, "\n"))
					}
				} else if errRead == io.EOF {
					if str != "" && str != "\n" {
						value = append(value, strings.Trim(str, "\n"))
					}
					break
				} else {
					fmt.Printf("Can't read file %s\n", elem)
					readError = true
					break
				}
			}
			if !readError {
				res[elem] = value
			}
			file.Close()
		} else {
			fmt.Printf("Can't open file %s\n", elem)
		}
	}
	return res
}

// Function find and return name file which has got string with match pattern
func fileNames(flags map[string]bool, pattern string, fileContent map[string][]string) []string {
	var addStr bool
	res := make([]string, 0)
	if flags["-i"] {
		pattern = strings.ToLower(pattern)
	}
	for file, txt := range fileContent {
		addStr = false
		for _, str := range txt {
			if flags["-i"] {
				str = strings.ToLower(str)
			}
			if flags["-v"] {
				addStr = reverseCompare(str, pattern, flags["-x"])
			} else {
				addStr = justCompare(str, pattern, flags["-x"])
			}
			if addStr {
				break
			}
		}
		if addStr {
			res = append(res, file)
		}
	}
	return res
}

// Assemble slice of strings witn match with pattern
func strMatch(flags map[string]bool, pattern string, fileContent map[string][]string, multiFile bool) []string {
	var addStr bool
	var sourceStr string

	res := make([]string, 0)
	if flags["-i"] {
		pattern = strings.ToLower(pattern)
	}
	for file, txt := range fileContent {
		addStr = false
		for strNum, str := range txt {
			sourceStr = str
			if flags["-i"] {
				str = strings.ToLower(str)
			}
			if flags["-v"] {
				addStr = reverseCompare(str, pattern, flags["-x"])
			} else {
				addStr = justCompare(str, pattern, flags["-x"])
			}
			if addStr {
				str = compileStr(sourceStr, strNum, file, multiFile, flags["-n"])
				res = append(res, str)
			}
		}
	}
	return res
}

// Construct string subject to exist many files and flag "-n"
func compileStr(str string, strNum int, file string, MultiFile bool, nFlag bool) string {
	if nFlag {
		str = strings.Join([]string{strconv.Itoa(strNum + 1), ":", str}, "")
	}
	if MultiFile {
		str = strings.Join([]string{file, ":", str}, "")
	}
	return str
}

// Function for parsing flags from os.Stdin
func parseFlags(arg string, i *int) ([]string, error) {
	if arg[0] != '-' {
		return make([]string, 0), nil
	}
	var res []string
	lenArg := len(arg)
	for i := 1; i < lenArg; i++ {
		if !validationFlag(flags, arg[i]) {
			err := strings.Join([]string{"Invalid flag [-", string(arg[i]), "]"}, "")
			return nil, errors.New(err)
		}
		switch arg[i] {
		case 'x':
			res = append(res, "-x")
		case 'l':
			res = append(res, "-l")
		case 'i':
			res = append(res, "-i")
		case 'n':
			res = append(res, "-n")
		case 'v':
			res = append(res, "-v")
		}
	}
	*i++
	return res, nil
}

func reverseCompare(str string, pattern string, wholeStr bool) bool {
	if wholeStr {
		if str != pattern {
			return true
		}
	} else {
		if !strings.Contains(str, pattern) {
			return true
		}
	}
	return false
}

func justCompare(str string, pattern string, wholeStr bool) bool {
	if wholeStr {
		if str == pattern {
			return true
		}
	} else {
		if strings.Contains(str, pattern) {
			return true
		}
	}
	return false
}

func validationFlag(flags [5]byte, opt byte) bool {
	for _, flag := range flags {
		if opt == flag {
			return true
		}
	}
	return false
}

func Usage() {
	fmt.Println("USAGE:")
	fmt.Printf("\tgo run grep [-flags] -pattern -files\n")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Printf("\t[-n] Print the line numbers of each matching line.\n")
	fmt.Printf("\t[-l] Print only the names of files that contain at least one matching line.\n")
	fmt.Printf("\t[-i] Match line using a case-insensitive comparison.\n")
	fmt.Printf("\t[-v] Invert the program -- collect all lines that fail to match the pattern.\n")
	fmt.Printf("\t[-x] Only match entire lines, instead of lines that contain a match.\n")
	fmt.Println()
	fmt.Println("Pattern")
	fmt.Printf("\t String which you need to find\n")
	fmt.Println()
	fmt.Println("Files:")
	fmt.Printf("\t List of files where need to find pattern\n")
}
