package modules

import (
	"golang.org/x/term"
	"strings"

	"bufio"
	"os"
	"syscall"
)

func ReadUserInputE() string {
	return ReadUserInput(true)
}

func ReadUserInput(echo bool) string {
	reader := bufio.NewReader(os.Stdin)
	input := ""
	if echo {
		input, _ = reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		input = strings.Replace(input, "\r", "", -1)
	} else {
		bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
		input = string(bytePassword)
	}
	return strings.TrimSpace(input)
}
