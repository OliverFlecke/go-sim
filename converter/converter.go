package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

func main() {
	filename := os.Args[1]

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
		return
	}

	var str strings.Builder
	var obj strings.Builder
	var x, y int
	for _, c := range string(content) {
		switch c {
		case '\n':
			y += 1
			x = -1
			str.WriteRune('\n')
		case ' ':
			str.WriteRune(' ')
		case '+':
			str.WriteRune('#')

		}
		switch {
		case 'a' <= c && c <= 'z':
			obj.WriteString(fmt.Sprintf("goal %c %d,%d\n", c, x, y))
			str.WriteRune(' ')
		case 'A' <= c && c <= 'Z':
			obj.WriteString(fmt.Sprintf("box %c %d,%d\n", unicode.ToLower(c), x, y))
			str.WriteRune(' ')
		case '0' <= c && c <= '9':
			obj.WriteString(fmt.Sprintf("agent %c %d,%d\n", c, x, y))
			str.WriteRune(' ')
		}
		x += 1
	}

	fmt.Print(str.String())
	fmt.Print("\n")
	fmt.Print(obj.String())
}
