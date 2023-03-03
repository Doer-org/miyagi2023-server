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

type CouponStatus struct {
	usecase usecase.CouponStatus
}

func NewCouponStatus(repository *registry.Repository) *CouponStatus {
	usecase := interactor.NewCouponStatus(
		repository.NewCouponStatus(),
	)
	return &CouponStatus{usecase: usecase}
}

// GET /coupon_statuses/{id}
func (h *CouponStatus) Get(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		response.BadRequestErr(w, fmt.Errorf("id param is empty"))
		return
	}

	in := &usecase.CouponStatusGetInput{}

	out, err := h.usecase.Get(r.Context(), in)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	view := newCouponStatusDefaultResponse(out.CouponStatus)
	response.New(w, view)
}

// POST /coupon_statues
func (h *CouponStatus) Create(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		response.BadRequestErr(w, fmt.Errorf("error: content length is 0"))
		return
	}

	var j couponStatusCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		response.BadRequestErr(w, err)
		return
	}

	in := &usecase.CouponStatusCreateInput{}

	out, err := h.usecase.Create(r.Context(), in)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	view := newCouponStatusDefaultResponse(out.CouponStatus)
	response.New(w, view)
}

// PUT /coupon_statuses/{id}/users/{user_id}/use
func (h *CouponStatus) Use(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		response.BadRequestErr(w, fmt.Errorf("id param is empty"))
		return
	}

	userIDParam := chi.URLParam(r, "user_id")
	if userIDParam == "" {
		response.BadRequestErr(w, fmt.Errorf("id param is empty"))
		return
	}

	in := &usecase.CouponStatusUseInput{}

	out, err := h.usecase.Use(r.Context(), in)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	view := newCouponStatusUseResponse(out.CouponStatus)
	response.New(w, view)
}

// GET /coupon_statues/list/users/{user_id}
func (h *CouponStatus) ListByUserID(w http.ResponseWriter, r *http.Request) {
	userIDParam := chi.URLParam(r, "user_id")
	if userIDParam == "" {
		response.BadRequestErr(w, fmt.Errorf("user_id param is empty"))
		return
	}

	in := &usecase.CouponStatusListByUserInput{}

	out, err := h.usecase.ListByUserID(r.Context(), in)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	view := newCouponStatusListByUserIDResponse(out.CouponStatuses)
	response.New(w, view)
}

type couponStatusDefaultResponse struct {
	ID        string                 `json:"id"`
	UsedFlg   bool                   `json:"used_flg"`
	CreatedAt string                 `json:"created_at"`
	UpdatedAt string                 `json:"updated_at"`
	Coupon    *couponDefaultResponse `json:"coupon"`
}

type couponStatusCreateRequest struct {
	CouponID string `json:"coupon_id"`
	UserID   string `json:"user_id"` //MEMO: 本来はsessionとかからとるべき?
}

// FIXME: 名前がいまいち
// 意図としてはUserのデータを使わず、ユーザーIDのみを返したいが
// List(),Use()で共通して使うためそのような名前にしたかった
type couponStatusWithUserIDResponse struct {
	ID        string                 `json:"id"`
	UsedFlg   bool                   `json:"used_flg"`
	CreatedAt string                 `json:"created_at"`
	UpdatedAt string                 `json:"updated_at"`
	Coupon    *couponDefaultResponse `json:"coupon"`
	User      string                 `json:"user_id"`
}

type couponStatusListByUserIDResponse struct {
	CouponStatues []*couponStatusWithUserIDResponse `json:"coupon_statuses"`
}

type couponStatusUseResponse struct {
	CouponStatus *couponStatusWithUserIDResponse `json:"coupon_status"`
}

func newCouponStatusDefaultResponse(couponStatus *model.CouponStatus) *couponStatusDefaultResponse {
	return &couponStatusDefaultResponse{
		ID:        couponStatus.ID.String(),
		UsedFlg:   couponStatus.UsedFlg,
		CreatedAt: couponStatus.CreatedAt.String(),
		UpdatedAt: couponStatus.UpdatedAt.String(),
		Coupon:    newCouponDefaultResponse(couponStatus.Coupon),
	}
}

func newCouponStatusesDefaultResponse(couponStatues []*model.CouponStatus) []*couponStatusDefaultResponse {
	var r []*couponStatusDefaultResponse
	for _, couponStatus := range couponStatues {
		r = append(r, newCouponStatusDefaultResponse(couponStatus))
	}
	return r
}

func newCouponStatusListByUserIDResponse(couponStatues []*model.CouponStatus) *couponStatusListByUserIDResponse {
	var lst []*couponStatusWithUserIDResponse
	for _, couponStatus := range couponStatues {
		lst = append(lst, newCouponStatusWithUserIDResponse(couponStatus))
	}
	return &couponStatusListByUserIDResponse{
		CouponStatues: lst,
	}
}

func newCouponStatusWithUserIDResponse(couponStatus *model.CouponStatus) *couponStatusWithUserIDResponse {
	return &couponStatusWithUserIDResponse{
		ID:        couponStatus.ID.String(),
		UsedFlg:   couponStatus.UsedFlg,
		CreatedAt: couponStatus.CreatedAt.String(),
		UpdatedAt: couponStatus.UpdatedAt.String(),
		Coupon:    newCouponDefaultResponse(couponStatus.Coupon),
		User:      couponStatus.User.ID.String(),
	}
}

func newCouponStatusUseResponse(couponStatus *model.CouponStatus) *couponStatusWithUserIDResponse {
	return newCouponStatusWithUserIDResponse(couponStatus)
}
