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

	user, err := uc.repository.Get(ctx, in.ID)
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

	// TODO: エラーハンドリング
	found,_ := uc.repository.Get(ctx,in.ID)
	if found != nil {
		return nil,fmt.Errorf("user id is exists")
	}

	// バリデーション
	if in.Name == "" {
		return nil,fmt.Errorf("user name is invalid")
	}

	if in.Age == 0 {
		return nil,fmt.Errorf("user age is invalid")
	}

	var gender model.Gender
	switch in.Gender {
	case string(model.MEN) :
		gender = model.Gender(in.Gender)
	case string(model.WOMEN):
		gender = model.Gender(in.Gender)
	default:
		return nil,fmt.Errorf("user gender is invalid")
	}

	birthday,err := time.Parse("2006-01-02",in.Birthday)
	if err != nil {
		return nil,err
	}

	if in.Address == "" {
		return nil,fmt.Errorf("user address is invalid")
	}

	u := &model.User{
		ID: model.UserID(in.ID),
		Name: in.Name,
		Age: in.Age,
		Gender: gender,
		Birthday: birthday,
		Address: in.Address,
		ProfileImg: in.ProfileImg, //必須ではない
		Prefecture: model.Prefecture(in.Prefecture), //FIXME: 必須?
	}

	newUser,err := uc.repository.Create(ctx,u)
	if err != nil {
		return nil,err
	}

	return &usecase.UserCreateOutput{
		User: newUser,
	},nil
}
