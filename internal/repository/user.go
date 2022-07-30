package repository

import (
	"context"

	"github.com/maxnosib/cognition/internal/entity"
)

func (m Model) CreateUser(ctx context.Context, usr entity.User) (int, error) {
	err := m.QueryRow("INSERT INTO users (nik, password) VALUES ($1,$2) RETURNING id", usr.Nik, usr.Pwd).Scan(&usr.ID)
	return usr.ID, err
}

func (m Model) GetByNik(ctx context.Context, nik string) (entity.User, error) {
	var usr entity.User
	err := m.QueryRow("SELECT * FROM users WHERE nik=$1", nik).Scan(&usr.ID, &usr.Nik, &usr.Pwd)

	return usr, err
}
