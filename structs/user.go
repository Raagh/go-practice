package structs

type User struct {
	userId    string
	firstName string
	lastName  string
	age       int
}

func (u *User) GetOlder() {
	u.age++
}

func (u User) GetFullName() string {
	// test := User{
	// 	userId: "c807143a-ceec-4838-9371-bcdaa4b83e8f",
	// 	age:    0,
	// }
	//
	// name := test.GetFullName()
	//
	// /*
	// * it should be one year older after this
	//  */
	// test.GetOlder()

	return u.firstName + " " + u.lastName
}
