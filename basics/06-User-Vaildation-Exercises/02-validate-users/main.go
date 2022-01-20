package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [usename] [password]"
	errUser  = "Access denied for %q\n"
	errPwd   = "Invalid password for %q\n"
	accessOK = "Access granted to %q.\n"

	user = "Berker"
	pass = "0000"

	user2 = "Berke"
	pass2 = "0001"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(usage)
		return
	}
	u, p := os.Args[1], os.Args[2]
	if u != user && u != user2 {
		fmt.Printf(errUser, u)
	} else if (p == pass && u == user) || (p == pass2 && u == user2) {
		fmt.Printf(accessOK, u)
	} else {
		fmt.Printf(errPwd, u)
	}
}
