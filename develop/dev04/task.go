package main

import (
	"fmt"
	"hash/fnv"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func getByteSliceHash(bytes []byte) uint32 {
	h := fnv.New32a()
	h.Write(bytes)
	return h.Sum32()
}

func parseAndSortStringToBytes(word string) []byte {
	bytes := []byte(word)
	sort.Slice(
		bytes,
		func(i, j int) bool {
			return (bytes)[i] < (bytes)[j]
		})
	return bytes
}

func getMapWithHashKeyWords(words []string) map[uint32][]string {
	res := make(map[uint32][]string, len(words))
	var hash uint32

	for _, v := range words {
		hash = getByteSliceHash(parseAndSortStringToBytes(v))
		if _, found := res[hash]; !found {
			res[hash] = make([]string, 0)
			res[hash] = append(res[hash], v)
			continue
		}
		res[hash] = append(res[hash], v)
	}
	return res
}

func makeStringsLowercase(words []string) []string {
	for i, v := range words {
		words[i] = strings.ToLower(v)
	}
	return words
}

func getConvertedMap(hashMap map[uint32][]string) map[string][]string {
	res := make(map[string][]string, len(hashMap))
	var t []string

	for _, v := range hashMap {
		if len(v) <= 1 {
			continue
		}
		t = v[1:]
		t = makeStringsLowercase(t)
		sort.Strings(t)
		res[v[0]] = t
	}
	return res
}

func main() {
	a := []string{"пятак", "листок", "слиток", "пятка", "тяпка", "столик", "кошка", "шокка"}
	r := getConvertedMap(getMapWithHashKeyWords(a))
	fmt.Println(r)
}
