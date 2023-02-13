package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.
Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const escapeSymbol = 92 // Unicode номер обратной косой черты

func StringUnpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	if _, err := strconv.Atoi(string(s[0])); err == nil {
		return "", errors.New("incorrect string")
	}

	runes := []rune(s)
	var builder strings.Builder
	var runeToPrint rune
	for i := 0; i < len(runes); i++ {

		if runes[i] == escapeSymbol { // если встретилась escape - последовательность, то
			i++ // переходим к следующему символу
		}
		runeToPrint = runes[i] // тот символ, который на очереди для печати

		if i+1 != len(runes) { // проверка на выход за пределы среза
			if unicode.IsDigit(runes[i+1]) {
				num, err := strconv.Atoi(string(runes[i+1]))
				if err != nil {
					return "", errors.New("can not convert int to string")
				}
				for j := 0; j < num; j++ { // добавляем в стринг билдер столько символов, сколько составляет интовое значение после этого символа
					builder.WriteRune(runeToPrint)
				}
				i++
			} else { // в противном случае просто добавляем символ в билдер
				builder.WriteRune(runeToPrint)
			}
		} else { // запись последнего символа
			builder.WriteRune(runeToPrint)
		}
	}

	return builder.String(), nil
}
