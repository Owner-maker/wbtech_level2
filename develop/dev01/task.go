package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

/*
=== Базовая задача ===

Создать программу, печатающую точное время с использованием NTP библиотеки. Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу, печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

const hostName = "0.beevik-ntp.pool.ntp.org"

func getHostTime(host string) (time.Time, error) {
	hostTime, err := ntp.Time(host)
	if err != nil {
		return time.Time{}, err
	}
	return hostTime, nil
}

func main() {
	t, err := getHostTime(hostName)
	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		if err != nil {
			os.Exit(1)
		}
		os.Exit(1)
	}
	fmt.Print(t)
}
