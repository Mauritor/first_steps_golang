package main

import "fmt"

type User struct {
	age            int
	name, lastName string
}

func main() {
	//mauritor := User{age: 45, name: "Mauri", lastName: "Tor"}
	//mauritor := User{45, "Mauri", "Tor"}

	mauritor := new(User)
	mauritor.name = "Mauri"
	fmt.Println(mauritor)
	fmt.Println(mauritor.name)
}
