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
)

type StampLog struct {
	usecase usecase.StampLog
}

func NewStampLog(repositroy *registry.Repository) *StampLog {
	usecase := interactor.NewStampLog(
		repositroy.NewStampLog(),
	)
	return &StampLog{
		usecase: usecase,
	}
}

func (h *StampLog) Create(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		response.BadRequestErr(w, fmt.Errorf("error: content length is 0"))
		return
	}

	var j stampLogCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		response.BadRequestErr(w, err)
		return
	}

	in := &usecase.StampLogCreateInput{}
	out, err := h.usecase.Create(r.Context(), in)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	view := newStampLogDefaultResponse(out.StampLog)
	response.New(w, view)
}

func (h *StampLog) List(w http.ResponseWriter, r *http.Request) {
	in := &usecase.StampLogListInput{}

	out, err := h.usecase.List(r.Context(), in)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	view := newStampLogListResponse(out.StampLogs)
	response.New(w, view)
}

type stampLogCreateRequest struct {
	SpotID      string `json:"spot_id"`
	UserID      string `json:"user_id"`
	StampCardID string `json:"stamp_card_id"` //必須ではない
}

type stampLogDefaultResponse struct {
	ID        string                    `json:"id"`
	Spot      *spotDefaultResponse      `json:"spot"`
	User      *userDefaultResponse      `json:"user"`
	StampCard *stampCardDefaultResponse `json:"stamp_card"`
	CreatedAt string                    `json:"created_at"`
}

type stampLogListResponse struct {
	StampLogs []*stampLogDefaultResponse `json:"stamp_logs"`
}

func newStampLogDefaultResponse(stampLog *model.StampLog) *stampLogDefaultResponse {
	return &stampLogDefaultResponse{
		ID:        stampLog.ID.String(),
		Spot:      newSpotDefaultResponse(stampLog.Spot),
		User:      newUserDefaultResponse(stampLog.User),
		StampCard: newStampCardDefaultResponse(stampLog.StampCard),
		CreatedAt: stampLog.CreatedAt.String(),
	}
}

func newStampLogsDefaultResponse(stampLogs []*model.StampLog) []*stampLogDefaultResponse {
	var r []*stampLogDefaultResponse
	for _, stampLog := range stampLogs {
		r = append(r, newStampLogDefaultResponse(stampLog))
	}
	return r
}

func newStampLogListResponse(stampLogs []*model.StampLog) *stampLogListResponse {
	return &stampLogListResponse{
		StampLogs: newStampLogsDefaultResponse(stampLogs),
	}
}
