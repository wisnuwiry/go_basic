package cli

import (
	"fmt"

	"go-trial-class/config"
	"go-trial-class/entity"
	"go-trial-class/helpers"
)

func Login() {
	helpers.ClearScreen()

	var user entity.User

	username := Input("Masukkan Username:")
	password := Input("Masukkan Password:")

	err := config.DB.Where("Username = ? AND Password = ?", username, password).First(&user).Error
	if err != nil {
		fmt.Println("Username/password salah")
		BackToMainMenu()
		return
	}

	user.SayHello()
}
