package proxy

import "fmt"

type UsersStack []User

func (u *UsersStack) Find(ID int) (User, error) {
	for _, user := range *u {
		if user.ID == ID {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user not found")
}

func (u *UsersStack) Add(user User) *UsersStack {
	*u = append(*u, user)
	return u
}

type UserFinderProxy struct {
	MainDB   UsersDB
	Stack    UsersStack
	Capacity int
}

func (u *UserFinderProxy) Find(ID int) (User, error) {
	user, err := u.Stack.Find(ID)
	if err == nil {
		fmt.Println("Found in stack: ", user)
		return user, nil
	}

	user, err = u.MainDB.Find(ID)
	if err != nil {
		return User{}, err
	}

	fmt.Println("Found in mainDB: ", user)
	u.AddToStack(user)
	return user, nil
}

func (u *UserFinderProxy) AddToStack(user User) error {
	fmt.Println("Adding to stack: ", user)
	if len(u.Stack) >= u.Capacity {
		u.Stack = append(u.Stack[1:], user)
	} else {
		u.Stack.Add(user)
	}
	return nil
}
