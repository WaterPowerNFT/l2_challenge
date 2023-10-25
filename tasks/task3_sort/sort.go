package main

import (
	"bufio"
	"flag"
	"fmt"
	_ "fmt"
	"os"
	"sort"
	"strings"
)

type doubleStr struct {
	field     string
	allString string
}

type launchFlags struct {
	u, r, n bool
	k       int
	f       string

	massive []doubleStr
}

func (curFlags *launchFlags) GetStrings() {
	file, err := os.Open(curFlags.f)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		newLine := new(doubleStr)
		newLine.allString = scanner.Text()
		if curFlags.k == 0 {
			newLine.field = scanner.Text()
		} else {
			splittedStr := strings.Split(scanner.Text(), " ")
			if len(splittedStr) <= curFlags.k {
				newLine.field = splittedStr[0]
			} else {
				newLine.field = splittedStr[curFlags.k]
			}

		}
		curFlags.massive = append(curFlags.massive, *newLine)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

func (curFlags *launchFlags) GetFlags() {

	f := flag.String("f", "", "specify file to sort")
	k := flag.Int("k", 0, "specify column")
	r := flag.Bool("r", false, "reverse order")
	u := flag.Bool("u", false, "dont print duplicates")
	n := flag.Bool("n", false, "sort by numeric")
	flag.Parse()

	if *f == "" {
		panic("Not set file")
	}
	curFlags.n = *n
	curFlags.f = *f
	curFlags.k = *k
	curFlags.r = *r
	curFlags.u = *u
}

func (curFlags *launchFlags) SortStrings() {
	if curFlags.r {
		sort.Slice(curFlags.massive, func(i, j int) (less bool) {
			return curFlags.massive[i].field > curFlags.massive[j].field
		})
	} else {
		sort.Slice(curFlags.massive, func(i, j int) (less bool) {
			return curFlags.massive[i].field < curFlags.massive[j].field
		})
	}
}

func (curFlags launchFlags) PrintStrings() {
	fmt.Println(curFlags.massive[0].allString)
	for i := 1; i < len(curFlags.massive); i++ {
		if curFlags.u && curFlags.massive[i].allString == curFlags.massive[i-1].allString {
			continue
		}
		fmt.Println(curFlags.massive[i].allString)
	}
}

func main() {
	var varCollector launchFlags
	varCollector.GetFlags()
	varCollector.GetStrings()
	varCollector.SortStrings()
	varCollector.PrintStrings()
}
