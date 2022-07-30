package usecase

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/maxnosib/cognition/internal/entity"
)

func (uc Model) CreateUser(ctx context.Context, user entity.User) (int, error) {
	hash := sha256.Sum256([]byte(user.Pwd))
	user.Pwd = fmt.Sprintf("%x", hash)

	id, err := uc.ur.CreateUser(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("CreateUser create err: %w", err)
	}

	return id, nil
}

func (uc Model) LoginUser(ctx context.Context, data entity.User) (int, error) {
	hash := sha256.Sum256([]byte(data.Pwd))
	data.Pwd = fmt.Sprintf("%x", hash)

	usr, err := uc.ur.GetByNik(ctx, data.Nik)
	if err != nil {
		return 0, fmt.Errorf("LoginUser get by id err: %w", err)
	}

	if usr.Pwd != data.Pwd {
		return 0, fmt.Errorf("invalid password")
	}

	return usr.ID, nil
}
