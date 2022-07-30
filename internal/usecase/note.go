package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/maxnosib/cognition/internal/entity"
)

func (uc Model) GetNotes(ctx context.Context, userID int) ([]entity.Note, error) {
	return uc.nr.GetNotes(ctx, userID)
}

func (uc Model) GetNoteByID(ctx context.Context, id string, userID int) (entity.Note, error) {
	return uc.nr.GetNoteByID(ctx, id, userID)
}

func (uc Model) CreateNote(ctx context.Context, note entity.Note, userID int) error {
	note.UserID = userID
	note.CreatedAt = time.Now().UTC()
	note.LastUpdatedAt = time.Now().UTC()
	if _, ok := uc.Categories[note.Category]; !ok {
		return fmt.Errorf("create note err: no category")
	}

	ok, err := uc.checkTag(ctx, note.Tags, userID)
	if !ok || err != nil {
		return fmt.Errorf("create note err: no tag")
	}

	ok, err = uc.checkLinks(ctx, note.Links, userID)
	if !ok || err != nil {
		return fmt.Errorf("create note err: invalid links")
	}

	rowNum, err := uc.nr.CreateNote(ctx, note)
	if rowNum == 0 || err != nil {
		return fmt.Errorf("create note err: %w", err)
	}

	return nil
}

func (uc Model) UpdateNote(ctx context.Context, note entity.Note, userID int) error {
	note.UserID = userID
	note.LastUpdatedAt = time.Now().UTC()

	if _, ok := uc.Categories[note.Category]; !ok {
		return fmt.Errorf("create note err: no category")
	}

	ok, err := uc.checkTag(ctx, note.Tags, userID)
	if !ok || err != nil {
		return fmt.Errorf("create note err: invalid tag")
	}

	ok, err = uc.checkLinks(ctx, note.Links, userID)
	if !ok || err != nil {
		return fmt.Errorf("create note err: invalid links")
	}

	rowNum, err := uc.nr.UpdateNote(ctx, note)
	if rowNum == 0 || err != nil {
		return fmt.Errorf("create note err: %w", err)
	}

	return nil
}

func (uc Model) GetCategories(ctx context.Context) (entity.Categories, error) {
	arrCat, err := uc.nr.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	res := make(entity.Categories)
	for _, val := range arrCat {
		res[val] = struct{}{}
	}

	return res, nil
}

func (uc Model) checkTag(ctx context.Context, tags []string, userID int) (bool, error) {
	if len(tags) == 0 {
		return false, fmt.Errorf("no tag")
	}

	data, err := uc.tr.GetTagsByTag(ctx, tags, userID)
	if err != nil {
		return false, err
	}

	mapTag := make(map[string]struct{})
	for _, val := range data {
		mapTag[val.Tag] = struct{}{}
	}

	for _, tag := range tags {
		if _, ok := mapTag[tag]; !ok {
			return false, fmt.Errorf("no tag")
		}
	}

	return true, nil
}

func (uc Model) checkLinks(ctx context.Context, links []string, userID int) (bool, error) {
	if len(links) == 0 {
		return true, nil
	}

	data, err := uc.nr.GetNotesByIDs(ctx, links, userID)
	if err != nil {
		return false, err
	}

	mapLinks := make(map[string]struct{})
	for _, val := range data {
		mapLinks[val.ID] = struct{}{}
	}

	for _, link := range links {
		if _, ok := mapLinks[link]; !ok {
			return false, fmt.Errorf("no link")
		}
	}

	return true, nil
}
