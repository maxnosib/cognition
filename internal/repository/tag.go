package repository

import (
	"context"

	"github.com/maxnosib/cognition/internal/entity"
)

func (m Model) GetTags(ctx context.Context, userID int) ([]entity.Tag, error) {
	var tags []entity.Tag

	rows, err := m.Query("SELECT * FROM tags WHERE user_id=$1", userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tag := entity.Tag{}
		err := rows.Scan(&tag.Tag, &tag.Description, &tag.UserID)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, rows.Err()
}

func (m Model) GetTagsByTag(ctx context.Context, tag []string, userID int) ([]entity.Tag, error) {
	var tags []entity.Tag

	rows, err := m.Query("SELECT * FROM tags WHERE user_id=$1 AND tag = any($2)", userID, tag)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tag := entity.Tag{}
		err := rows.Scan(&tag.Tag, &tag.Description, &tag.UserID)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, rows.Err()
}

func (m Model) CreateTag(ctx context.Context, tag entity.Tag) (int64, error) {
	rows, err := m.Exec("INSERT INTO tags (tag, description, user_id) VALUES ($1,$2,$3)", tag.Tag, tag.Description, tag.UserID)
	return rows.RowsAffected(), err
}

func (m Model) DeleteTag(ctx context.Context, tag string, userID int) (int64, error) {
	rows, err := m.Exec("DELETE FROM tags WHERE tag=$1 and user_id=$2", tag, userID)
	return rows.RowsAffected(), err
}
