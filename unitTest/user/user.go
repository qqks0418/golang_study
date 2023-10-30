package user

import "fmt"

type User struct {
	ID    int64
	Name  string
	Email string
}

func NewUser(id int64, name string, email string) *User {
	if id == 1 {
		fmt.Println("A")
	} else {
		fmt.Println("B")
	}

	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}