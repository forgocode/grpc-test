package server

import (
	"fmt"
	"serverMonitor/pkg/typed"
)

func Login(user typed.User) bool {
	dbUser := GetUserByName(user.Name)
	if dbUser == nil {
		return false
	}
	if dbUser.Passwd != user.Passwd {
		return false
	}
	return true
}

func Regiter(user typed.User) error {
	if GetUserByName(user.Name) != nil {
		return fmt.Errorf("user has been exist")
	}
	if !checkName(user.Name) {
		return fmt.Errorf("error user name")
	}
	if !checkPasswd(user.Passwd) {
		return fmt.Errorf("error user passwd")
	}
	return AddUserToDb(&user)

}

func GetUserByName(name string) *typed.User {
	return nil
}

func AddUserToDb(user *typed.User) error {
	return nil
}

func checkName(name string) bool {
	return true
}

func checkPasswd(passwd string) bool {
	return true
}
