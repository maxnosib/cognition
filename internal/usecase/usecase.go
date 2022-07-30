package usecase

import (
	"context"

	"github.com/maxnosib/cognition/internal/entity"
)

var Categories entity.Categories

type userRepository interface {
	CreateUser(ctx context.Context, usr entity.User) (int, error)
	GetByNik(ctx context.Context, id string) (entity.User, error)
}
type tagRepository interface {
	GetTags(ctx context.Context, userID int) ([]entity.Tag, error)
	GetTagsByTag(ctx context.Context, tags []string, userID int) ([]entity.Tag, error)
	CreateTag(ctx context.Context, tag entity.Tag) (int64, error)
	DeleteTag(ctx context.Context, tag string, userID int) (int64, error)
}
type noteRepository interface {
	GetNotes(ctx context.Context, userID int) ([]entity.Note, error)
	GetNotesByIDs(ctx context.Context, ids []string, userID int) ([]entity.Note, error)
	GetNoteByID(ctx context.Context, id string, userID int) (entity.Note, error)
	CreateNote(ctx context.Context, note entity.Note) (int64, error)
	UpdateNote(ctx context.Context, note entity.Note) (int64, error)
	GetCategories(ctx context.Context) ([]string, error)
	DeleteTagInNote(ctx context.Context, tag string, userID int) (int64, error)
}

type Model struct {
	ur         userRepository
	tr         tagRepository
	nr         noteRepository
	Categories entity.Categories
}

func New(cat entity.Categories, ur userRepository, tr tagRepository, nr noteRepository) Model {
	return Model{ur: ur, tr: tr, nr: nr, Categories: cat}
}
