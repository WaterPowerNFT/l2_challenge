package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

type params struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	fileName   string
	regExp     string
}

func getOsParams() *params {
	toRet := &params{}
	flag.IntVar(&toRet.after, "A", 0, "Print +N rows after match")
	flag.IntVar(&toRet.before, "B", 0, "Print +N rows before match")
	flag.IntVar(&toRet.context, "C", 0, "Print +N rows after and before match")

	flag.BoolVar(&toRet.count, "c", false, "Print count of match rows")
	flag.BoolVar(&toRet.count, "i", false, "Ignore case")
	flag.BoolVar(&toRet.count, "v", false, "Instead of a match, exclude")
	flag.BoolVar(&toRet.count, "F", false, "Exact match with a string, not a pattern")
	flag.BoolVar(&toRet.count, "n", false, "Print line number of match rows")

	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		panic(fmt.Sprintf("Wrong num of args. Needed 2, got %v\n", len(args)))
	} else {
		toRet.fileName = args[1]
		toRet.regExp = args[0]
	}
	return toRet
}

func readFile(filename string) ([]string, error) {
	rows := []string{}
	file, err := os.Open(filename)
	if err != nil {
		return rows, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		rows = append(rows, sc.Text())
	}
	return rows, nil
}

func (p *params) grep() (interface{}, error) {
	fileRows, err := readFile(p.fileName)
	if err != nil {
		return 0, err
	}
	pref, post := "", ""
	if p.ignoreCase {
		pref = "(?i)"
	}

	if p.fixed {
		pref += "^"
		post = "$"
	}

	regularExp, errorExp := regexp.Compile(pref + p.regExp + post)
	if errorExp != nil {
		return 0, errorExp
	}

	switch {
	case p.after != 0:
		{
			for i, row := range fileRows {
				if regularExp.MatchString(row) {
					if p.after <= len(fileRows)-i {
						return fileRows[i : i+p.after+1], nil
					}

					return fileRows[i:], nil
				}
			}
			return "not found", nil
		}

	case p.before != 0:
		{
			for i, row := range fileRows {
				if regularExp.MatchString(row) {
					if p.before-1 <= i {
						return fileRows[i-p.before : i+1], nil
					}
					return fileRows[:i+1], nil
				}
			}
			return "not found", nil
		}

	case p.context != 0:
		{
			for i, row := range fileRows {
				if regularExp.MatchString(row) {
					startIndex := 0
					endIndex := len(fileRows)

					if p.context-1 <= i {
						startIndex = i - p.context
					}

					if p.context <= len(fileRows)-i {
						endIndex = i + p.context + 1
					}

					return fileRows[startIndex:endIndex], nil
				}
			}
			return "not found", nil
		}

	case p.count:
		{
			totalCount := 0
			for _, row := range fileRows {
				totalCount += len(regularExp.FindAllString(row, -1))
			}

			if p.invert {
				return len(fileRows) - totalCount, nil
			}

			return totalCount, nil
		}

	case p.lineNum:
		{
			numberOfRows := []int{}
			for i, row := range fileRows {
				if regularExp.MatchString(row) {
					numberOfRows = append(numberOfRows, i)
				}
			}
			return numberOfRows, nil
		}

	default:
		{
			result := []string{}
			for _, row := range fileRows {
				if regularExp.MatchString(row) {
					result = append(result, row)
				}
			}
			return result, nil
		}
	}
}

func main() {
	myParams := getOsParams()
	result, err := myParams.grep()
	if err != nil {
		panic(err)
	}

	switch typeResult := result.(type) {
	case string:
		{
			fmt.Println(result)
		}
	case []string:
		{
			for _, row := range typeResult {
				fmt.Println(row)
			}
		}
	case int:
		{
			fmt.Println(result)
		}
	case []int:
		{
			for _, row := range typeResult {
				fmt.Printf("%d ", row)
			}
		}
	default:
		{
			fmt.Printf("Unknown result. Type of result %T", result)
		}
	}
}
