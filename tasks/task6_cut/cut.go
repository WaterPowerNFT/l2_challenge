package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	fl "./fl"
)

func main() {
	launchFlags := fl.FlagConstructor()
	f := flag.String("f", "", "select only these fields; also print any line that contains no delimiter character, unless the -s option is specified")
	d := flag.String("d", "\t", "use DELIM instead of TAB for field delimite")
	s := flag.Bool("s", false, "do not print lines not containing delimiters")

	flag.Parse()
	//fmt.Println(*f, *d, *s)
	err := launchFlags.GetArgs(f, d, s)
	if err != nil {
		panic(err)
	}
	//fmt.Println("Lets specify strings. Print quit to stop input")
	reader := bufio.NewReader(os.Stdin)
	cameString, _ := reader.ReadString('\n')
	for cameString != "quit\r\n" {
		launchFlags.AddNewString(cameString)
		cameString, _ = reader.ReadString('\n')
	}
	fmt.Println(launchFlags.ReturnFinalresult())
}
