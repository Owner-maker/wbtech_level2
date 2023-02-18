package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const fileOutput = "output.txt"

func writeHtmlBodyToFile(file string, response http.Response) error {
	data, err := os.Create(file)
	if err != nil {
		return err
	}
	defer func(data *os.File) {
		err = data.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(data)

	writer := bufio.NewWriter(data)
	_, err = io.Copy(writer, response.Body)
	if err != nil {
		return err
	}
	return nil
}

func wget(file, url string) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	err = writeHtmlBodyToFile(file, *r)
	if err != nil {
		return err
	}
	return err
}

func main() {
	url := "https://owner-maker.github.io/"
	err := wget(fileOutput, url)
	if err != nil {
		log.Fatal(err)
	}
}
