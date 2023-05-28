package cli

import "fmt"

func ErrorHandler(msg string) {
	fmt.Println("Oops ada error:", msg)
}
