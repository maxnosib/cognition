package entity

import "fmt"

type User struct {
	ID  int
	Nik string
	Pwd string `db:"password"`
}

func (usr User) Validate(_ int) error {
	if usr.Nik == "" {
		return fmt.Errorf("nik is empty")
	}
	if usr.Pwd == "" {
		return fmt.Errorf("password is empty")
	}

	return nil
}
