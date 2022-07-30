package entity

import "fmt"

type Tag struct {
	Tag         string
	Description string
	UserID      int
}

type GetTag struct {
	Tag string
}

func (tg GetTag) Validate(_ int) error {
	if tg.Tag == "" {
		return fmt.Errorf("tag is empty")
	}

	return nil
}

func (tg *Tag) Validate(userID int) error {
	if tg.Tag == "" {
		return fmt.Errorf("tag is empty")
	}
	if tg.Description == "" {
		return fmt.Errorf("description is empty")
	}
	if tg.UserID == 0 {
		tg.UserID = userID
	}

	return nil
}
