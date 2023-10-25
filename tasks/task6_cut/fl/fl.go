package fl

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

// FlagsElems has fields with flags and final elems to print
type FlagsElems struct {
	_d string
	_f []int
	_s bool

	_elemsToPrint string
}

// ReturnFinalresult return elems to print
func (fe FlagsElems) ReturnFinalresult() string {
	return fe._elemsToPrint
}

// AddNewString adding new string to elemsToPrint
func (fe *FlagsElems) AddNewString(cameString string) {
	var sb strings.Builder
	massiveFields := strings.Split(cameString, fe._d)
	if (len(massiveFields) == 1 && fe._s == true) || (fe._f[0] != 0) {
		return
	}
	i := 0
	sb.WriteString(fe._elemsToPrint)
	for i < len(fe._f) && fe._f[i] < len(massiveFields) {
		sb.WriteString(massiveFields[fe._f[i]])
		sb.WriteRune(' ')
		i++
	}
	sb.WriteRune('\n')
	fe._elemsToPrint = sb.String()
}

// FlagConstructor return FlagsElems unit
func FlagConstructor() FlagsElems {
	var flagToReturn FlagsElems
	return flagToReturn
}

// SliceAtoi convert string slice to int slice
func SliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

// GetArgs get args, which came with launch
func (fe *FlagsElems) GetArgs(f, d *string, s *bool) error {

	if len(*f) < 1 {
		return errors.New("you need to specify fields")
	}
	intArr, err := SliceAtoi(strings.Split(*f, ","))
	if err != nil {
		return err
	}
	sort.Ints(intArr)
	fe._f = intArr
	fe._d = *d
	fe._s = *s
	return nil
}
