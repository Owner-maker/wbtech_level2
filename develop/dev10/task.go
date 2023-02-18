package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

const deliveryProtocol = "tcp"

type Connection struct {
	port    int
	host    string
	timeout time.Duration
}

func parseFlags() (Connection, error) {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		return Connection{}, errors.New("not enough arguments")
	}
	var host string
	var port int
	var timeout int

	for _, v := range args {
		r := strings.Split(v, "=")
		if len(r) == 1 {
			return Connection{}, fmt.Errorf("can not parse argument -> %s", r[0])
		}
		if r[0] == "host" {
			host = r[1]
			continue
		}
		if r[0] == "port" {
			v, err := strconv.Atoi(r[1])
			if err != nil {
				return Connection{}, fmt.Errorf("can not parse <port> argument -> %s", r[1])
			}
			port = v
			continue
		}
		if r[0] == "timeout" {
			v, err := strconv.Atoi(r[1])
			if err != nil {
				return Connection{}, fmt.Errorf("can not parse <timeout> argument -> %s", r[1])
			}
			timeout = v
			continue
		}
		return Connection{}, fmt.Errorf("unknown argument <%s>", r[0])
	}

	if host == "" {
		return Connection{}, errors.New("argument <host> is missed")
	}

	return Connection{
		port:    port,
		host:    host,
		timeout: time.Duration(timeout) * time.Second,
	}, nil
}

func consoleReader(c net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		t, err := reader.ReadString('\n')
		if err != nil {
			err = c.Close()
			if err != nil {
				log.Fatal(err)
			}
			log.Fatal("connection lost")
		}
		_, err = c.Write([]byte(t))
		if err != nil {
			c.Close()
			log.Fatalln(err)
		}
	}
}

func main() {
	flags, err := parseFlags()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	conn, err := net.DialTimeout(
		deliveryProtocol,
		net.JoinHostPort(flags.host, strconv.Itoa(flags.port)), flags.timeout)

	// попытка подключения
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn net.Conn) {
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	go consoleReader(conn)

	fmt.Println("connection to the server was successfully initiated")

	// канал для закрытия от пользователя
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGQUIT)

	<-sigCh
}
