package usecase

import (
	"context"
	"fmt"

	"github.com/maxnosib/cognition/internal/entity"
)

func (uc Model) GetTags(ctx context.Context, userID int) ([]entity.Tag, error) {
	return uc.tr.GetTags(ctx, userID)
}

func (uc Model) CreateTag(ctx context.Context, tag entity.Tag) error {
	rowNum, err := uc.tr.CreateTag(ctx, tag)
	if rowNum == 0 || err != nil {
		return fmt.Errorf("create tag err: %w", err)
	}

	return nil
}

func (uc Model) DeleteTag(ctx context.Context, tag string, userID int) error {
	rowNum, err := uc.tr.DeleteTag(ctx, tag, userID)
	if rowNum == 0 || err != nil {
		return fmt.Errorf("delete tag err: %w", err)
	}

	uc.nr.DeleteTagInNote(ctx, tag, userID)
	if rowNum == 0 || err != nil {
		return fmt.Errorf("delete tag err: %w", err)
	}

	return nil
}
