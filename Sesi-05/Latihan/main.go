package main

import (
	"fmt"
	"errors"

	// "github.com/gopasspw/gopass"
)

func main() {
	defer catchErr()

	var usrname string
	var passwd []byte

	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&usrname)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&passwd)

	// password, _ := gopass.GetPasswdMasked()

	if valid, err := validLogin(usrname , passwd); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error Ocured:", r)
	} else {
		fmt.Println("Aplikasi Berjalan Sempurna")
	}
}


func validLogin(username string, password []byte) (string, error) {
	s := string(password[:])

	if username == "abdul" && s == "123" {
		return "Berhasil Log in", nil
	} else {
		return "", errors.New("Username/Password Salah!") 
	}
}
