package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type ArrayFlags []string

func (i *ArrayFlags) String() string {
	return ""
}

func (i *ArrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var (
	arrayFlags    ArrayFlags
	delimiterFlag = flag.String("d", "\t", "use another delimiter")
	separatedFlag = flag.Bool("s", true, "only lines with delimiter")
)

func cut(lines []string) ([]string, error) {
	var res []string
	var builder strings.Builder

	for _, v := range lines {
		tempStrings := strings.Split(v, *delimiterFlag)
		if (*separatedFlag && len(tempStrings) > 1) || (!*separatedFlag) {
			for _, column := range arrayFlags {
				val, err := strconv.Atoi(column)
				if err != nil {
					return nil, err
				}

				if val <= len(tempStrings) {
					builder.WriteString(tempStrings[val-1])
					builder.WriteString(*delimiterFlag)
				}
			}
			builder.WriteString("\n")
			res = append(res, builder.String())
			builder.Reset()
		}
	}

	return res, nil
}

func main() {
	arrayFlags = []string{"1", "3"}
	flag.Var(&arrayFlags, "f", "select fields (columns)")
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	lines := make([]string, n)
	i := 0

	for range lines {
		str, err := bufio.NewReader(os.Stdin).ReadString('\r')
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		lines[i] = str
		i++
	}

	r, err := cut(lines)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println()
	for _, v := range r {
		fmt.Print(v)
	}
}
