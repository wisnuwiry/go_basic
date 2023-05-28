package cli

import (
	"bufio"
	"fmt"
	"os"
)

func Input(instruction string) string {
	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println(instruction)
	input, err := consoleReader.ReadString('\n')
	if err != nil {
		ErrorHandler(err.Error())
	}

	return input
}
