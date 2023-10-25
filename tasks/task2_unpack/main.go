package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func deployStr(strToDeploy string) (string, error) {
	if len(strToDeploy) == 0 {
		return "", nil
	}
	char := strToDeploy[0]
	if char >= 48 && char <= 57 {
		return "", errors.New("incorrect string")
	}
	var num int = 0
	var sb strings.Builder
	for i := 1; i < len(strToDeploy)-1; i++ {
		curChar := strToDeploy[i]
		if curChar >= 48 && curChar <= 57 {
			num *= 10
			num += int(curChar - 48)
		} else {
			sb.WriteString(strings.Repeat(string(char), max(num, 1)))
			char = curChar
			num = 0
		}
	}
	sb.WriteString(strings.Repeat(string(char), max(num, 1)))
	return sb.String(), nil
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	newText, err := deployStr(text)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newText)
	}
}
