package proxy

import "fmt"

type UsersDB []User

func (u *UsersDB) Find(ID int) (User, error) {
	for _, user := range *u {
		if user.ID == ID {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user not found")
}

func (u *UsersDB) Add(user User) *UsersDB {
	fmt.Println("Adding to database: ", user)
	*u = append(*u, user)
	return u
}
