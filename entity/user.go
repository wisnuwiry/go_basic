package entity

import "fmt"

type User struct {
	ID       int
	Username string
	Password string
}

func (u *User) SayHello() {
	fmt.Println("Selamat Datang ", u.Username)
}
