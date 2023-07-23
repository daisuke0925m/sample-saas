package domain

import "fmt"

type Users []User

type User struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	FullName  string `json:"full_name,omitempty"`
}

func (u *User) Build() *User {
	u.FullName = fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	return u
}
