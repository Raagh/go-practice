package main

import "fmt"

type User struct {
	id        string
	firstName string
	lastName  string
	age       int
}

func (u *User) GetOlder() {
	u.age++
}

func (u User) GetFullName() string {
	return u.firstName + " " + u.lastName
}

func main() {
	test := User{
		id:        "c807143a-ceec-4838-9371-bcdaa4b83e8f",
		firstName: "Patricio",
		lastName:  "Lopez",
		age:       0,
	}

	test.firstName = "Ramon"

	name := test.GetFullName()

	fmt.Println("This is the name: ", name)

	/*
	* it should be one year older after this
	 */
	test.GetOlder()

	fmt.Println("This is the age", test.age)
}
