package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const (
	inputFileName  = "input.txt"
	outputFileName = "output.txt"
)

var (
	columnSortFlag              = flag.Int("c", -1, "specifying a column to sort")
	numbersSortFlag             = flag.Bool("n", false, "sort by numeric value")
	reverseSortFlag             = flag.Bool("r", false, "sort in reverse order")
	withoutDuplicateStringsFlag = flag.Bool("u", false, "do not output duplicate lines")
)

func readLinesFromFile(file string) ([]string, error) {
	data, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	var res []string
	r := bufio.NewReader(data)
	const d = '\n'
	for {
		l, err := r.ReadString(d)
		if err == nil || len(l) > 0 {
			if err != nil {
				l += string(d)
			}
			res = append(res, l)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return res, nil
}

func writeLinesToFile(file string, strings []string) error {
	data, err := os.Create(file)
	if err != nil {
		return err
	}
	defer data.Close()

	if *reverseSortFlag {
		strings = reverse(strings)
	}
	if *withoutDuplicateStringsFlag {
		strings = excludeRepeatedStrings(strings)
	}

	for _, v := range strings {
		_, err = data.WriteString(fmt.Sprintf("%s", v))
		if err != nil {
			return err
		}
	}
	return nil
}

func sortByColumn(lines []string, columnInd int) ([]string, error) {
	var strByColumns [][]string
	var res []string

	for _, v := range lines {
		values := strings.Split(v, " ")
		if len(values)-1 < columnInd {
			return nil, fmt.Errorf("there is no such column = %d", columnInd)
		}
		strByColumns = append(strByColumns, values)
	}

	sort.Slice(strByColumns, func(i, j int) bool {
		if len(strByColumns[i]) == 0 || len(strByColumns[j]) == 0 {
			return len(strByColumns[i]) == 0
		}

		return strByColumns[i][columnInd] < strByColumns[j][columnInd]
	})
	for _, v := range strByColumns {
		res = append(res, strings.Join(v, " "))
	}

	return res, nil
}

func sortByNumbers(lines []string) ([]string, error) {
	var res []string
	var tempNums []int
	var tempStr string

	for _, v := range lines {
		tempStr = strings.TrimSuffix(v, "\n")
		tempStr = strings.Replace(tempStr, "\r", "", 1)
		n, err := strconv.Atoi(tempStr)

		if err != nil {
			return nil, err
		}
		tempNums = append(tempNums, n)
	}

	sort.Ints(tempNums)
	for _, v := range tempNums {
		res = append(res, fmt.Sprintf("%s\n", strconv.Itoa(v)))
	}
	return res, nil
}

func reverse(lines []string) []string {
	res := make([]string, len(lines))
	l := len(lines) - 1
	for i, _ := range lines {
		res = append(res, lines[l-i])
	}
	return res
}

func excludeRepeatedStrings(lines []string) []string {
	set := make(map[string]struct{})
	var res []string

	for _, v := range lines {
		if _, found := set[v]; !found {
			res = append(res, v)
			set[v] = struct{}{}
		}
	}
	return res
}

func main() {
	flag.Parse()

	r, err := readLinesFromFile(inputFileName)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if *numbersSortFlag {
		r, err = sortByNumbers(r)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	} else if *columnSortFlag >= 0 {
		r, err = sortByColumn(r, *columnSortFlag)
	}

	err = writeLinesToFile(outputFileName, r)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
