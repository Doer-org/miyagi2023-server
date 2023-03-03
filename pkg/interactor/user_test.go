package interactor

import (
	"context"
	"testing"

	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/Doer-org/miyagi2023-server/pkg/interactor/mock"
)

func TestGet(t *testing.T) {
	t.Skip()
	tests := []struct {
		name    string
		reqId   string
		wantErr bool
	}{
		{
			name:    "正常に動作した場合",
			reqId:   "xxx-xxxx-xx",
			wantErr: false,
		},
		{
			name:    "idが空なら error",
			reqId:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userusecase := NewUser(&UserRepositoryMock{})
		})
	}

}

type UserRepositoryMock struct{}

func (*UserRepositoryMock) Get(context.Context, model.UserID) (*model.User, error) {
	return mock.User, nil
}
func (*UserRepositoryMock) Create(context.Context, *model.User) (*model.User, error) {
	return mock.User, nil
}
