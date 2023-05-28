package cli

import (
	"fmt"
	"os"

	"go-trial-class/helpers"
)

func MainMenu() {
	helpers.ClearScreen()

	fmt.Println("------------------")
	fmt.Println("AUTHENTIATION APP")
	fmt.Println("------------------")

	fmt.Println(" 1. Login")
	fmt.Println(" 2. Register")
	fmt.Println(" q. Keluar Aplikasi")

	fmt.Println("------------------")

	fmt.Println("Silahkan Masukkan Menu:")

	var input string
	fmt.Scanln(&input)

	switch input {
	case "1":
		Login()
	case "2":
		Register()
	case "q":
		fmt.Println("Terimakasih telah mengunjungi aplikasi")
		os.Exit(1)
	default:
		MainMenu()
	}
}

func BackToMainMenu() {
	Input("Tekan enter untuk kembali ke main menu")

	MainMenu()
}
