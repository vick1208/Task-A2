package main

import "strings"

func ValidInput(firstName string, lastName string, email string, userTic uint, remTic uint) (bool, bool, bool) {
	isVName := len(firstName) >= 2 && len(lastName) >= 2
	isVEmail := strings.Contains(email, "@")
	isVTNum := userTic > 0 && userTic <= remTic
	return isVName, isVEmail, isVTNum
}
