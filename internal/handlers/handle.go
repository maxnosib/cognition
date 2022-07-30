package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/maxnosib/cognition/internal/entity"
	"github.com/maxnosib/cognition/internal/usecase"
	"github.com/maxnosib/cognition/utils"
)

func GetNote(ctx context.Context, request *http.Request, us usecase.Model, userID int) utils.Answer {
	var answer utils.Answer
	id := chi.URLParam(request, "id")
	answer.Data, answer.Err = us.GetNoteByID(ctx, id, userID)

	return answer
}

func GetNotes(ctx context.Context, request *http.Request, us usecase.Model, userID int) utils.Answer {
	var answer utils.Answer

	answer.Data, answer.Err = us.GetNotes(ctx, userID)

	return answer
}

func UpdateNote(ctx context.Context, request *http.Request, us usecase.Model, userID int) utils.Answer {
	var answer utils.Answer
	var dataIN entity.Note
	answer.Err = utils.ReadJSONAndValidate(request, &dataIN, userID)
	if answer.Err != nil {
		return answer
	}

	answer.Err = us.UpdateNote(ctx, dataIN, userID)

	return answer
}

func AddNote(ctx context.Context, request *http.Request, us usecase.Model, userID int) utils.Answer {
	var answer utils.Answer
	var dataIN entity.Note
	answer.Err = utils.ReadJSONAndValidate(request, &dataIN, userID)
	if answer.Err != nil {
		return answer
	}

	answer.Err = us.CreateNote(ctx, dataIN, userID)

	return answer
}

func Auth(ctx context.Context, r *http.Request, us usecase.Model) (entity.User, error) {
	var usr entity.User
	err := utils.ReadJSONAndValidate(r, &usr, 0)
	if err != nil {
		return usr, err
	}

	usr.ID, err = us.LoginUser(ctx, usr)

	return usr, err
}

func Register(ctx context.Context, r *http.Request, us usecase.Model) (entity.User, error) {
	var usr entity.User
	err := utils.ReadJSONAndValidate(r, &usr, 0)
	if err != nil {
		return usr, err
	}

	usr.ID, err = us.CreateUser(ctx, usr)

	return usr, err
}

func GetTags(ctx context.Context, request *http.Request, us usecase.Model, userID int) utils.Answer {
	var answer utils.Answer

	answer.Data, answer.Err = us.GetTags(ctx, userID)

	return answer
}

func AddTag(ctx context.Context, request *http.Request, us usecase.Model, userID int) utils.Answer {
	var answer utils.Answer
	var dataIN entity.Tag
	answer.Err = utils.ReadJSONAndValidate(request, &dataIN, userID)
	if answer.Err != nil {
		return answer
	}

	answer.Err = us.CreateTag(ctx, dataIN)

	return answer
}

func DeleteTag(ctx context.Context, request *http.Request, us usecase.Model, userID int) utils.Answer {
	var answer utils.Answer
	id := chi.URLParam(request, "id")

	answer.Err = us.DeleteTag(ctx, id, userID)

	return answer
}

func GetCategories(ctx context.Context, request *http.Request, us usecase.Model, userID int) utils.Answer {
	var answer utils.Answer
	cat, err := us.GetCategories(ctx)
	if err != nil {
		answer.Err = err
		return answer
	}

	res := make([]string, 0)
	for key := range cat {
		res = append(res, key)
	}

	answer.Data = res

	return answer
}
