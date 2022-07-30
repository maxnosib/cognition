package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID            string
	Category      string
	UserID        int
	Description   string
	Links         []string
	Tags          []string
	Sources       string
	IsFirst       bool
	CreatedAt     time.Time
	LastUpdatedAt time.Time
}

type GetNote struct {
	ID string
}

type Categories map[string]struct{}

func (nt GetNote) Validate(_ int) error {
	if nt.ID == "" {
		return fmt.Errorf("id is empty")
	}

	return nil
}

func (c Categories) CheckCategoryExists(cat string) error {
	if _, ok := c[cat]; ok {
		return nil
	}

	return fmt.Errorf("no category")
}

func (nt *Note) Validate(userID int) error {
	if nt.ID == "" {
		nt.ID = uuid.New().String()
	}
	if nt.Category == "" {
		return fmt.Errorf("category_id is empty")
	}
	if nt.UserID == 0 {
		nt.UserID = userID
	}
	if nt.Description == "" {
		return fmt.Errorf("description is empty")
	}
	if len(nt.Tags) == 0 {
		return fmt.Errorf("tags is empty")
	}

	return nil
}
