package main

import "fmt"

// when
// 1.when we need to update state
// pointer = 8 bytes
// 2. when we want to optimize the memory for Large objects that are gettiing called A Lot
type User struct {
	email    string
	username string
	age      int
	file     []byte // ?? Large ??
}

func getUser() (User, error) {
	return &User{}, fmt.Errorf("foo")
}

func (u User) Email() string {
	return u.email
}

// func (u *User) updateEmail(email string) {
// 	u.email = email
// }

func (u *User) UpdateEmail(email string) {
	u.email = email
}

// x amount of bytes => sizeOf (user)
// 1 gb user size
func Email(user User) string {
	return user.email
}

func main() {
	user := User{
		email: "afff@foo.com",
	}
	user.UpdateEmail("foo@bar.com")
	fmt.Println(user.Email())
}
