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

type Spot struct {
	usecase usecase.Spot
}

func NewSpot(r *registry.Repository) *Spot {
	usecase := interactor.NewSpot(
		r.NewSpot(),
	)
	return &Spot{
		usecase: usecase,
	}
}

// GET /spots/{id}
func (h *Spot) Get(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		response.BadRequestErr(w, fmt.Errorf("id param is empty"))
		return
	}

	in := &usecase.SpotGetInput{}
	out, err := h.usecase.Get(r.Context(), in)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	view := newSpotDefaultResponse(out.Spot)
	response.New(w, view)
}

// POST /spots
func (h *Spot) Create(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		response.BadRequestErr(w, fmt.Errorf("error: content length is 0"))
		return
	}

	var j spotCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		response.BadRequestErr(w, err)
		return
	}

	in := &usecase.SpotCreateInput{}
	out, err := h.usecase.Create(r.Context(), in)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	view := newSpotDefaultResponse(out.Spot)
	response.New(w, view)
}

// GET /spots/list
func (h *Spot) List(w http.ResponseWriter, r *http.Request) {
	in := &usecase.SpotListInput{}

	out, err := h.usecase.List(r.Context(), in)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	view := newSpotListResponse(out.Spots)
	response.New(w, view)
}

type spotDefaultResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Detail    string `json:"detail"`
	Like      uint   `json:"like"`
	ImgURL    string `json:"img_url"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
}

type spotCreateRequest struct {
	Name    string `json:"name"`
	Detail  string `json:"detail"`
	ImgURL  string `json:"img_url"`
	Address string `json:"address"`
}

type spotListResponse struct {
	Spots []*spotDefaultResponse `json:"spots"`
}

func newSpotDefaultResponse(spot *model.Spot) *spotDefaultResponse {
	return &spotDefaultResponse{
		ID:        spot.ID.String(),
		Name:      spot.Name,
		Detail:    spot.Detail,
		Like:      spot.Like,
		ImgURL:    spot.ImgURL,
		Address:   spot.Address,
		CreatedAt: spot.CreatedAt.String(),
	}
}

func newSpotsDefaultResponse(spots []*model.Spot) []*spotDefaultResponse {
	var r []*spotDefaultResponse
	for _, spot := range spots {
		r = append(r, newSpotDefaultResponse(spot))
	}
	return r
}

func newSpotListResponse(spots []*model.Spot) *spotListResponse {
	return &spotListResponse{
		Spots: newSpotsDefaultResponse(spots),
	}
}
