package user

import (
	"errors"
	"fmt"
)

type User struct {
	Name  string
	Money float64
	Items map[string]int
}

func CreateUser(name string, money float64) User {
	new_user := new(User)
	new_user.Items = make(map[string]int)
	new_user.Money = money
	new_user.Name = name
	return *new_user
}

func (this User) PrintUserItems() {
	fmt.Println(this.Items)
}

func (this User) CheckBalance(cost float64) error {
	if this.Money < cost {
		return errors.New("Not enough money")
	}
	return nil
}
