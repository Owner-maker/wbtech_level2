package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func executeCommand(command string) error {
	command = strings.Replace(command, "\r", "", 1)
	command = strings.Replace(command, "\n", "", 1)
	commands := strings.Split(command, " ")
	l := len(commands)

	if l < 1 {
		return errors.New("invalid command")
	}

	switch commands[0] {
	case "exit":
		os.Exit(0)
	case "cd":
		if l != 2 {
			return errors.New("command <cd> accept only 1 argument -> path")
		}
		moveToDirectory(commands[1])
		break
	case "pwd":
		if l != 1 {
			return errors.New("command <pwd> do not accept any arguments")
		}
		fmt.Println(getCurrentPath())
	case "echo":
		if l != 2 {
			return errors.New("command <echo> accepts only 1 argument -> command name")
		}
		fmt.Println(commands[1])
	case "kill":
		if l != 2 {
			return errors.New("command <kill> accepts only 1 argument -> process id")
		}
		id, err := strconv.Atoi(commands[1])
		if err != nil {
			return errors.New("invalid argument")
		}
		err = findAndKillProcess(id)
		if err != nil {
			return err
		}
	case "ps":
		if l != 1 {
			return errors.New("command <ps> do not accept any arguments")
		}
		list, err := getListOfRunningProcesses()
		if err != nil {
			return err
		}
		for _, v := range list {
			fmt.Print(v)
		}
	default:
		return fmt.Errorf("unknown command <%s>", commands[0])
	}
	return nil
}

func getListOfRunningProcesses() ([]string, error) {
	processList, err := ps.Processes()
	if err != nil {
		return nil, err
	}
	var process ps.Process
	s := make([]string, len(processList))

	for x := range processList {
		process = processList[x]
		s[x] = fmt.Sprintf("%d\t%s\n", process.Pid(), process.Executable())
	}
	return s, nil
}

func findAndKillProcess(id int) error {
	p, err := os.FindProcess(id)
	if err != nil {
		return err
	}
	err = p.Kill()
	if err != nil {
		return err
	}
	return nil
}

func getCurrentPath() string {
	curPath, err := os.Getwd()
	if err != nil {
		return ""
	}
	return curPath
}

func moveToDirectory(path string) {
	err := os.Chdir(path)
	if err != nil {
		return
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">>>")
		command, err := reader.ReadString('\r')
		if err != nil {
			_, err = fmt.Fprintln(os.Stderr, err.Error())
			if err != nil {
				os.Exit(1)
			}
		}

		err = executeCommand(command)
		if err != nil {
			_, err = fmt.Fprintln(os.Stderr, err.Error())
			if err != nil {
				os.Exit(1)
			}
		}
	}
}
