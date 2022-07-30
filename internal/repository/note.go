package repository

import (
	"context"
	"strings"

	"github.com/maxnosib/cognition/internal/entity"
)

func (m Model) GetNotes(ctx context.Context, userID int) ([]entity.Note, error) {
	var notes []entity.Note

	rows, err := m.Query("SELECT id, category, user_id, description, links, tags, sources, is_first, created_at, last_updated_at FROM notes WHERE user_id=$1", userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		note := entity.Note{}
		err := rows.Scan(&note.ID, &note.Category, &note.UserID, &note.Description, &note.Links, &note.Tags, &note.Sources, &note.IsFirst, &note.CreatedAt, &note.LastUpdatedAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, rows.Err()
}

func (m Model) GetNoteByID(ctx context.Context, id string, userID int) (entity.Note, error) {
	var note entity.Note
	err := m.QueryRow("SELECT id, category, user_id, description, links, tags, sources, is_first, created_at, last_updated_at FROM notes WHERE id=$1 AND user_id=$2", id, userID).
		Scan(&note.ID, &note.Category, &note.UserID, &note.Description, &note.Links, &note.Tags, &note.Sources, &note.IsFirst, &note.CreatedAt, &note.LastUpdatedAt)
	return note, err
}

func (m Model) CreateNote(ctx context.Context, note entity.Note) (int64, error) {
	rows, err := m.Exec("INSERT INTO notes (id, category, user_id, description, links, tags, sources, is_first, created_at, last_updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)",
		&note.ID, &note.Category, &note.UserID, &note.Description, &note.Links, &note.Tags, &note.Sources, &note.IsFirst, &note.CreatedAt, &note.LastUpdatedAt)
	return rows.RowsAffected(), err
}

func (m Model) UpdateNote(ctx context.Context, note entity.Note) (int64, error) {
	rows, err := m.Exec("UPDATE notes SET category=$1, description=$2, links=$3, tags=$4, sources=$5, is_first=$6, last_updated_at=$7 WHERE id=$8 and user_id=$9",
		note.Category, note.Description, note.Links, note.Tags, note.Sources, note.IsFirst, note.LastUpdatedAt, note.ID, note.UserID)
	return rows.RowsAffected(), err
}

func (m Model) DeleteTagInNote(ctx context.Context, tag string, userID int) (int64, error) {
	rows, err := m.Exec("UPDATE notes SET tags = array_remove(tags, $1) WHERE user_id=$2;", tag, userID)
	return rows.RowsAffected(), err
}

func (m Model) GetCategories(ctx context.Context) ([]string, error) {
	var res []string
	err := m.QueryRow("SELECT enum_range(NULL::categories)").Scan(&res)
	return res, err
}

func (m Model) GetNotesByIDs(ctx context.Context, ids []string, userID int) ([]entity.Note, error) {
	var notes []entity.Note

	rows, err := m.Query("SELECT id, category, user_id, description, links, tags, sources, is_first, created_at, last_updated_at FROM notes WHERE user_id=$1  AND tag IN($2)", userID, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		note := entity.Note{}
		err := rows.Scan(&note.ID, &note.Category, &note.UserID, &note.Description, &note.Links, &note.Tags, &note.Sources, &note.IsFirst, &note.CreatedAt, &note.LastUpdatedAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, rows.Err()
}
