package main

import (
	"fmt"
	"strings"
)

func URLfy(s string) string {

	lenOriginal := len(s)
	numEspacos := strings.Count(s, " ")
	newLength := lenOriginal + 3*numEspacos
	newString := make([]byte, newLength)

	j := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			newString[j] = '%'
			newString[j+1] = '2'
			newString[j+2] = '0'
			j += 3
		} else {
			newString[j] = s[i]
			j++
		}
	}

	return string(newString)
}

func main() {
	var test = "elison monteiro passos "
	fmt.Printf("%v\n", URLfy(test))
}
