package cli

import (
	"go-trial-class/config"
	"go-trial-class/entity"
	"go-trial-class/helpers"
)

func Register() {
	helpers.ClearScreen()

	username := Input("Masukkan Username:")
	password := Input("Masukkan Password:")

	user := entity.User{
		Username: username,
		Password: password,
	}

	err := config.DB.Create(&user).Error
	if err != nil {
		ErrorHandler(err.Error())
		BackToMainMenu()
		return
	}

	user.SayHello()
}
