package cipher

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type caesarCipher struct {
	shift rune
}

type vigenereCipher struct {
	codeString string
}

func Run() {
	var res string
	if len(os.Args) != 2 || (os.Args[1] != "-cd" && os.Args[1] != "-vd" &&
		os.Args[1] != "-ce" && os.Args[1] != "-ve") {
		Usage()
		return
	}
	sc := bufio.NewScanner(os.Stdin)
	if os.Args[1][1] == 'c' {
		fmt.Printf("Input shift: ")
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		str := sc.Text()
		shift, err := strconv.Atoi(str)
		if err != nil || (!(shift > -26 && shift < 0) && !(shift > 0 && shift < 26)) {
			fmt.Println("Error: Shift must be -25 <= x <= -1 or 1 <= x <= 25")
			return
		}
		Caesar := NewCaesar(shift)
		fmt.Println("Input phrase: ")
		sc.Scan()
		phrase := sc.Text()
		if os.Args[1] == "-ce" {
			res = Caesar.Encode(phrase)
			fmt.Println("Encode text: ")
		} else if os.Args[1] == "-cd" {
			res = Caesar.Decode(phrase)
			fmt.Println("Decode text: ")
		}
	} else if os.Args[1][1] == 'v' {
		fmt.Printf("Input code phrase: \n")
		sc.Scan()
		codePhrase := sc.Text()
		Vigenere, err := NewVigenere(codePhrase)
		if err != nil {
			//Usage()
			fmt.Println(err)
			return
		}
		fmt.Println("Input phrase: ")
		sc.Scan()
		phrase := sc.Text()
		if os.Args[1] == "-ve" {
			res = Vigenere.Encode(phrase)
			fmt.Println("Encode text: ")
		} else if os.Args[1] == "-vd" {
			res = Vigenere.Decode(phrase)
			fmt.Println("Decode text: ")
		}
	}
	fmt.Println(res)
}

func NewCaesar(shift int) Cipher {
	if (shift >= 1 && shift <= 25) || (shift <= -1 && shift >= -25) {
		var res = caesarCipher{shift: rune(shift)}
		return res
	}
	return nil
}

func NewVigenere(code string) (Cipher, error) {
	aWhole := true
	codeR := []rune(code)
	if len(codeR) == 0 {
		return nil, errors.New("Error: Code phrase is empty")
	}
	for _, r := range codeR {
		if r < 'a' || r > 'z' {
			return nil, errors.New("Invalid simbol in code phrase")
		} else if r != 'a' {
			aWhole = false
		}
	}
	if aWhole {
		return nil, errors.New("Error: Code phrase consist only from letter 'a'")
	}
	res := vigenereCipher{code}
	return res, nil
}

func (s caesarCipher) Encode(str string) string {
	str = strings.ToLower(str)
	strR := []rune(str)
	var res []rune
	for _, r := range strR {
		if r >= 'a' && r <= 'z' {
			if s.shift > 0 {
				res = append(res, 'a'+((r-'a'+s.shift)%26))
			} else {
				var k rune
				if r-'a' < -1*s.shift {
					k = 26 + s.shift
				} else {
					k = s.shift
				}
				res = append(res, r+k)
			}
		}
	}
	return string(res)
}

func (s caesarCipher) Decode(str string) string {
	str = strings.ToLower(str)
	strR := []rune(str)
	var res []rune
	for _, r := range strR {
		if r >= 'a' && r <= 'z' {
			if s.shift > 0 {
				var k rune
				if r-'a' < s.shift {
					k = 26 - s.shift
				} else {
					k = -1 * s.shift
				}
				res = append(res, r+k)
			} else {
				res = append(res, 'a'+((r-'a'-s.shift)%26))
			}
		}
	}
	return string(res)
}

func (v vigenereCipher) Encode(str string) string {
	var j int
	var res []rune
	str = strings.ToLower(str)
	strR := []rune(str)
	codeR := []rune(v.codeString)
	lenCode := len(codeR)
	for i := 0; i < len(strR); i++ {
		if strR[i] >= 'a' && strR[i] <= 'z' {
			res = append(res, 'a'+((strR[i]-'a')+(codeR[j]-'a'))%26)
			j++
			if j == lenCode {
				j = 0
			}
		}
	}
	return string(res)
}

func (v vigenereCipher) Decode(str string) string {
	var j int
	var res []rune
	str = strings.ToLower(str)
	strR := []rune(str)
	codeR := []rune(v.codeString)
	lenCode := len(codeR)
	for i := 0; i < len(strR); i++ {
		if strR[i] >= 'a' && strR[i] <= 'z' {
			res = append(res, 'z'-(('z'-strR[i])+(codeR[j]-'a'))%26)
			j++
			if j == lenCode {
				j = 0
			}
		}
	}
	return string(res)
}

func Usage() {
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("\tgo run cipher.go -mode")
	fmt.Println("Existing mode:")
	fmt.Println("\t[-ce] - Caesar encoding mode")
	fmt.Println("\t[-cd] - Caesar decoding mode")
	fmt.Println("\t[-ve] - Vigenere encoding mode")
	fmt.Println("\t[-vd] - Vigenere decoding mode")
	fmt.Println("\t[-cd] - Caesar decoding mode")
	fmt.Println()
	fmt.Println("In Caesar mode will be ofered to input shift. Shift can be -25 <= shift <= -1, or 1 <= shift <= 25")
	fmt.Println("In Vigenere mode will be ofered to input code phrase. Code phrase can't to be empty, ")
	fmt.Println("contain something other than english lowercase letter or consist only from letter 'a'")
	fmt.Println("In either case will be ofered to input phrase for encode/decode")

}
