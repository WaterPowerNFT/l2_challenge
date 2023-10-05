package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func DeployStr(str_to_deploy string) (string, error) {
	char := str_to_deploy[0]
	if char >= 48 && char <= 57 {
		return "", errors.New("incorrect string")
	}
	var num int = 0
	var sb strings.Builder
	for i := 1; i < len(str_to_deploy)-1; i += 1 {
		cur_char := str_to_deploy[i]
		if cur_char >= 48 && cur_char <= 57 {
			num *= 10
			num += int(cur_char - 48)
		} else {
			sb.WriteString(strings.Repeat(string(char), max(num, 1)))
			char = cur_char
			num = 0
		}
	}
	return sb.String(), nil
}
func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		new_text, err := DeployStr(text)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(new_text)
		}
	}
}
