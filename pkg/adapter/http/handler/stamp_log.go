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

func NewStampLog(repository *registry.Repository) *StampLog {
	usecase := interactor.NewStampLog(
		repository.NewStampLog(),
		repository.NewCoupon(),
		repository.NewCouponStatus(),
	)
	return &StampLog{
		usecase: usecase,
	}
}

// POST /stamp_logs
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

	in := &usecase.StampLogCreateInput{
		SpotID: j.SpotID,
		UserID: j.UserID,
	}
	out, err := h.usecase.Create(r.Context(), in)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	view := newStampLogCreateResponse(out.StampLog,out.VisitCnt)
	response.New(w, view)
}

// GET /stamp_logs/list
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
	SpotID string `json:"spot_id"`
	UserID string `json:"user_id"`
}

type stampLogCreateResponse struct {
	ID        string                 `json:"id"`
	Spot      *spotDefaultResponse   `json:"spot"`
	User      *userDefaultResponse   `json:"user"`
	Coupon    *couponDefaultResponse `json:"coupon"`
	CreatedAt string                 `json:"created_at"`
	VisitCnt uint `json:"visit_cnt"`
}

type _stampLogListResponse struct {
	ID        string `json:"id"`
	SpotID    string `json:"spot_id"`
	UserID    string `json:"user_id"`
	CouponID  string `json:"coupon_id"`
	CreatedAt string `json:"created_at"`
}

type stampLogListResponse struct {
	StampLogs []*_stampLogListResponse `json:"stamp_logs"`
}

func newStampLogCreateResponse(stampLog *model.StampLog,visitCnt uint) *stampLogCreateResponse {
	return &stampLogCreateResponse{
		ID:        stampLog.ID.String(),
		Spot:      newSpotDefaultResponse(stampLog.Spot),
		User:      newUserDefaultResponse(stampLog.User),
		CreatedAt: stampLog.CreatedAt.String(),
		VisitCnt: visitCnt,
	}
}


func _newStampLogListResponse(stampLogs []*model.StampLog) []*_stampLogListResponse {
	var r []*_stampLogListResponse
	for _, stampLog := range stampLogs {
		r = append(r, &_stampLogListResponse{
			ID:        stampLog.ID.String(),
			SpotID:    stampLog.Spot.ID.String(),
			UserID:    stampLog.User.ID.String(),
			CreatedAt: stampLog.CreatedAt.String(),
		})
	}
	return r
}

func newStampLogListResponse(stampLogs []*model.StampLog) *stampLogListResponse {
	return &stampLogListResponse{
		StampLogs: _newStampLogListResponse(stampLogs),
	}
}
