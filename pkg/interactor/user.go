package interactor

import (
	"context"
	"fmt"
	"time"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/Doer-org/miyagi2023-server/pkg/domain/repository"
	"github.com/Doer-org/miyagi2023-server/pkg/usecase"
)

type User struct {
	repository repository.User
}

func NewUser(repository repository.User) *User {
	return &User{repository: repository}
}

func (uc *User) Get(ctx context.Context, in *usecase.UserGetInput) (*usecase.UserGetOutput, error) {
	if in.ID == "" {
		return nil, fmt.Errorf("user id is invalid")
	}

	user, err := uc.repository.Get(ctx, model.UserID(in.ID))
	if err != nil {
		return nil, err
	}

	return &usecase.UserGetOutput{
		User: user,
	}, nil
}

func (uc *User) Create(ctx context.Context, in *usecase.UserCreateInput) (*usecase.UserCreateOutput, error) {
	if in.ID == "" {
		return nil, fmt.Errorf("user id is invalid")
	}
	if len(in.ID) > 255 {
		return nil, fmt.Errorf("user id must be 255 characters or less")
	}

	found, _ := uc.repository.Get(ctx, model.UserID(in.ID))
	if found != nil {
		return nil, fmt.Errorf("user id is exists")
	}

	if in.Name == "" {
		return nil, fmt.Errorf("user name is invalid")
	}
	if len(in.Name) > 255 {
		return nil, fmt.Errorf("user name must be 255 characters or less")
	}

	if in.Age == 0 {
		return nil, fmt.Errorf("user age is invalid")
	}
	if in.Age > 100 {
		return nil, fmt.Errorf("user age max is 100")
	}

	var gender model.Gender
	switch in.Gender {
	case string(model.MEN):
		gender = model.Gender(in.Gender)
	case string(model.WOMEN):
		gender = model.Gender(in.Gender)
	default:
		return nil, fmt.Errorf("user gender is invalid")
	}

	birthday, err := time.Parse("2006-01-02", in.Birthday)
	if err != nil {
		return nil, err
	}

	if in.Address == "" {
		return nil, fmt.Errorf("user address is invalid")
	}
	if len(in.Address) > 255 {
		return nil, fmt.Errorf("user address must be 255 characters or less")
	}

	u := &model.User{
		ID:         model.UserID(in.ID),
		Name:       in.Name,
		Age:        in.Age,
		Gender:     gender,
		Birthday:   birthday,
		Address:    in.Address,
		ImgURL:     in.ImgURL,                       //必須ではない
		Prefecture: model.Prefecture(in.Prefecture), //FIXME: 必須?
	}

	newUser, err := uc.repository.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	return &usecase.UserCreateOutput{
		User: newUser,
	}, nil
}
