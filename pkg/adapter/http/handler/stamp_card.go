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

type StampCard struct {
	usecase usecase.StampCard
}

func NewStampCard(repostiroy *registry.Repository) *StampCard {
	usecase := interactor.NewStampCard(
		repostiroy.NewStampCard(),
	)
	return &StampCard{
		usecase: usecase,
	}
}

func (h *StampCard) Get(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		response.BadRequestErr(w, fmt.Errorf("id param is empty"))
		return
	}

	in := &usecase.StampCardGetInput{}
	out, err := h.usecase.Get(r.Context(), in)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	view := newStampCardDefaultResponse(out.StampCard)
	response.New(w, view)
}

func (h *StampCard) Create(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		response.BadRequestErr(w, fmt.Errorf("error: content length is 0"))
		return
	}

	var j stampCardCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		response.BadRequestErr(w, err)
		return
	}

	in := &usecase.StampCardCreateInput{}
	out, err := h.usecase.Create(r.Context(), in)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	view := newStampCardDefaultResponse(out.StampCard)
	response.New(w, view)
}

func (h *StampCard) List(w http.ResponseWriter, r *http.Request) {
	in := &usecase.StampCardListInput{}

	out, err := h.usecase.List(r.Context(), in)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	view := newStampCardListResponse(out.StampCards)
	response.New(w, view)
}

type stampCardCreateRequest struct {
	Name string `json:"name"`
}
type stampCardDefaultResponse struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	CreatedAt string                 `json:"created_at"`
	Spots     []*spotDefaultResponse `json:"spots"`
}

type stampCardListResponse struct {
	StampCards []*stampCardDefaultResponse `json:"stamp_cards"`
}

func newStampCardDefaultResponse(stampCard *model.StampCard) *stampCardDefaultResponse {
	return &stampCardDefaultResponse{
		ID:        stampCard.ID.String(),
		Name:      stampCard.Name,
		CreatedAt: stampCard.CreatedAt.String(),
		Spots:     newSpotsDefaultResponse(stampCard.Spots),
	}
}

func newStampCardsDefaultResponse(stampCards []*model.StampCard) []*stampCardDefaultResponse {
	var r []*stampCardDefaultResponse
	for _, stampCard := range stampCards {
		r = append(r, newStampCardDefaultResponse(stampCard))
	}
	return r
}

func newStampCardListResponse(stampCards []*model.StampCard) *stampCardListResponse {
	return &stampCardListResponse{
		StampCards: newStampCardsDefaultResponse(stampCards),
	}
}
