package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func convertToAnagramas(mapKeySorted map[string][]string) *map[string][]string {
	dictAnagramas := make(map[string][]string)
	for _, value := range mapKeySorted {
		if len(value) > 1 {
			for i := 0; i < len(value); i++ {
				dictAnagramas[value[0]] = append(dictAnagramas[value[0]], value[i])
			}
			sort.Strings(dictAnagramas[value[0]])
		}
	}
	return &dictAnagramas
}

func getAnagramas(massiveWords *[]string) *map[string][]string {
	dictKeySorted := map[string][]string{}
	for _, elem := range *massiveWords {
		elem = strings.ToLower(elem)
		sortedElem := sortString(elem)
		found := slices.Contains(dictKeySorted[sortedElem], elem)
		if !found {
			dictKeySorted[sortedElem] = append(dictKeySorted[sortedElem], elem)
		}
	}
	return convertToAnagramas(dictKeySorted)
}

func main() {
	massive := []string{"авс", "пятак", "пятка", "тяпка", "листок", "слиток", "столик", "вас", "сав", "сва", "один", "сва"}
	someVar := getAnagramas(&massive)
	fmt.Println(someVar)
}
