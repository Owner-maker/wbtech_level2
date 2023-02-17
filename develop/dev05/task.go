package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const inputFileName = "input.txt"

var (
	afterFlag      = flag.Int("A", 0, "print +N lines after match")
	beforeFlag     = flag.Int("B", 0, "print +N lines until match")
	contextFlag    = flag.Int("C", 0, "print ±N lines around the match")
	countFlag      = flag.Bool("c", false, "quantity of lines")
	ignoreCaseFlag = flag.Bool("i", false, "ignore case")
	invertFlag     = flag.Bool("v", false, "exclude")
	fixedFlag      = flag.Bool("F", false, "exact string match, not a pattern")
	lineNumFlag    = flag.Bool("n", false, "print a number of line")
)

type Pair struct {
	Key   int
	Value string
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Key < p[j].Key }

func trimStrOfEndSymbols(s string) string {
	s = strings.TrimSuffix(s, "\n")
	s = strings.Replace(s, "\r", "", 1)
	return s
}

func addLinesBeforeMatchedOne(lines []string, matchInd, n int, m *map[int]string) {
	curInd := matchInd
	if n > 0 {
		border := matchInd - n
		if border < 0 {
			border = 0
		}
		for _, v := range lines[border:matchInd] {
			(*m)[curInd] = v
			curInd--
		}
	}
}

func addLinesAfterMatchedOne(lines []string, matchInd, n int, m *map[int]string) {
	var leftBorder int
	var rightBorder int
	curInd := matchInd + 2

	l := len(lines)

	if n > 0 {
		leftBorder = matchInd + 1
		rightBorder = matchInd + n + 1
		if leftBorder >= l {
			leftBorder = l - 1
		}
		if rightBorder >= l {
			rightBorder = l - 1
		}

		for _, v := range lines[leftBorder:rightBorder] {
			(*m)[curInd] = v
			curInd++
		}
	}
}

func convertMapValuesToStringSlice(m *map[int]string) []string {
	res := make([]string, len(*m))
	pairs := make(PairList, len(*m))
	var tempStr string

	i := 0
	for k, v := range *m {
		pairs[i] = Pair{k, v}
		i++
	}
	i = 0
	sort.Sort(pairs)
	for _, v := range pairs {
		tempStr = v.Value
		if *lineNumFlag {
			tempStr = fmt.Sprintf("%d %s", v.Key, v.Value)
		}

		res[i] = tempStr
		i++
	}
	return res
}

func grep(lines []string, pattern string) (int, []string, error) {
	var res []string
	m := make(map[int]string, len(lines))
	var tempStr string

	if *fixedFlag {
		for i, v := range lines {
			tempStr = trimStrOfEndSymbols(v)

			if *ignoreCaseFlag {
				tempStr = strings.ToLower(tempStr)
				pattern = strings.ToLower(tempStr)
			}

			if (tempStr == pattern && !*invertFlag) || (tempStr != pattern && *invertFlag) {
				m[i+1] = tempStr
				if *contextFlag > 0 {
					addLinesBeforeMatchedOne(lines, i, *contextFlag, &m)
					addLinesAfterMatchedOne(lines, i, *contextFlag, &m)
				} else {
					if *beforeFlag > 0 {
						addLinesBeforeMatchedOne(lines, i, *beforeFlag, &m)
					}
					if *afterFlag > 0 {
						addLinesAfterMatchedOne(lines, i, *afterFlag, &m)
					}
				}
			}
		}
		res = convertMapValuesToStringSlice(&m)
		return len(res), res, nil
	}

	if *ignoreCaseFlag {
		pattern = fmt.Sprintf("(?i)%s", pattern)
	}

	r, err := regexp.Compile(pattern)
	if err != nil {
		return 0, nil, err
	}

	for i, v := range lines {
		tempStr = trimStrOfEndSymbols(v)
		if (r.MatchString(tempStr) && !*invertFlag) || (!r.MatchString(tempStr) && *invertFlag) {
			m[i+1] = tempStr
			if *contextFlag > 0 {
				addLinesBeforeMatchedOne(lines, i, *contextFlag, &m)
				addLinesAfterMatchedOne(lines, i, *contextFlag, &m)
			} else {
				if *beforeFlag > 0 {
					addLinesBeforeMatchedOne(lines, i, *beforeFlag, &m)
				}
				if *afterFlag > 0 {
					addLinesAfterMatchedOne(lines, i, *afterFlag, &m)
				}
			}
		}
	}

	if *countFlag {
		return len(m), nil, nil
	}
	res = convertMapValuesToStringSlice(&m)
	return len(res), res, nil
}

func main() {
	flag.Parse()

	r := []string{"hi", "hello", "hey", "hello", "dude", "dude"}

	count, lines, err := grep(r, ".Y")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if len(lines) != 0 {
		for _, v := range lines {
			fmt.Println(v)
		}
	} else {
		fmt.Print(count)
	}

}
