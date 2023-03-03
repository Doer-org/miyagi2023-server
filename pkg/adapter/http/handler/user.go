package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Doer-org/miyagi2023-server/pkg/adapter/http/response"
	"github.com/Doer-org/miyagi2023-server/pkg/domain/model"
	"github.com/Doer-org/miyagi2023-server/pkg/infra/registry"
	"github.com/Doer-org/miyagi2023-server/pkg/interactor"
	"github.com/Doer-org/miyagi2023-server/pkg/usecase"
	"github.com/go-chi/chi"
)

type User struct {
	usecase usecase.User
}

func NewUser(r *registry.Repository) *User {
	usecase := interactor.NewUser(
		r.NewUser(),
	)
	return &User{usecase: usecase}
}

// GET /users/{id}
func (h *User) Get(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		response.BadRequestErr(w, fmt.Errorf("id param is empty"))
		return
	}

	in := &usecase.UserGetInput{
		ID: idParam,
	}

	out, err := h.usecase.Get(r.Context(), in)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	view := newUserDefaultResponse(out.User)
	response.New(w, view)
}

// POST /users
func (h *User) Create(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		response.BadRequestErr(w, fmt.Errorf("error: content length is 0"))
		return
	}

	var j userCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		response.BadRequestErr(w, err)
		return
	}

	in := &usecase.UserCreateInput{
		ID:         j.ID,
		Name:       j.Name,
		Age:        j.Age,
		Gender:     j.Gender,
		Birthday:   j.Birthday,
		Address:    j.Address,
		ProfileImg: j.ProfileImg,
		Prefecture: j.Prefecture,
	}

	out, err := h.usecase.Create(r.Context(), in)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	view := newUserDefaultResponse(out.User)
	response.New(w, view)
}

type userCreateRequest struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Age        uint   `json:"age"`
	Gender     string `json:"gender"`
	Birthday   string `json:"birthday"`
	Address    string `json:"address"`
	ProfileImg string `json:"profile_img"`
	Prefecture string `json:"prefecture"`
}

type userDefaultResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Age        uint   `json:"age"`
	Gender     string `json:"gender"`
	Birthday   string `json:"birthday"`
	Address    string `json:"address"`
	ProfileImg string `json:"profile_img"`
	Prefecture string `json:"prefecture"`
	CreatedAt  string `json:"created_at"`
}

func newUserDefaultResponse(user *model.User) *userDefaultResponse {
	return &userDefaultResponse{
		ID:         user.ID.String(),
		Name:       user.Name,
		Age:        user.Age,
		Gender:     user.Gender.String(),
		Birthday:   user.Birthday.String(),
		Address:    user.Address,
		ProfileImg: user.ProfileImg,
		Prefecture: user.Prefecture.String(),
		CreatedAt:  user.CreatedAt.String(),
	}
}
