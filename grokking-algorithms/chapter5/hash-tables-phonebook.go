package chapter5

import "errors"

// a map is a hash table in golang
//
// every string gets passed through a hashfunction
// gets converted into a number, and added into an array using the number as index
// every time you pass the string, that string gets converted to a number
// the number is used as index retrieving the stored value related to the string
var phonebook = make(map[string]int)

func SavePhonebookContact(name string, number int) {
	phonebook[name] = number
}

func GetPhoneNumber(name string) (int, error) {
	number, exists := phonebook[name]
	if exists {
		return number, nil
	}

	return -1, errors.New("couldnt find the number for this person")
}
