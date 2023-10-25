package main

import (
	"bufio"
	"os"
	"os/exec"
	"strings"

	currentInfo "./current_info"
	"golang.org/x/exp/errors/fmt"
)

func main() {
	additionalString := currentInfo.ConstructorCurInfo()
	for {
		additionalString.PrintNewScreen()
		prevData := ""
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			splittedStringsFromScanner := strings.Split(scanner.Text(), "|")
			for _, elem := range splittedStringsFromScanner {
				cameArgs := strings.Split(elem, " ")
				cameArgs = append(cameArgs, prevData)
				fmt.Println(cameArgs)
				if cameArgs[0] == "pwd" {
					additionalString.Echo(additionalString.GetPath())
				} else if cameArgs[0] == "echo" {
					if len(scanner.Text()) > 5 {
						additionalString.Echo(scanner.Text()[5:])
					}
				} else if cameArgs[0] == "ps" {
					additionalString.PS()
				} else if cameArgs[0] == "kill" {
					if len(cameArgs) == 2 {
						additionalString.KillProcessByName(cameArgs[1])
					} else {
						additionalString.Echo("wrong num of args for kill")
					}
				} else if cameArgs[0] == "cd" {
					if cameArgs[1] == ".." {
						additionalString.TryRemoveDir()
					} else {
						additionalString.TryAddDir(cameArgs[1])
					}
				} else {
					out, err := exec.Command(cameArgs[0]).Output()
					if err != nil {
						fmt.Println(err)
					}
					prevData = string(out)
					//fmt.Println(prevData)
				}

			}
		} else {
			panic("scanner trouble")
		}

	}
}
