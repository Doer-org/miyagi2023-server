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

type Coupon struct {
	usecase usecase.Coupon
}

func NewCoupon(repository *registry.Repository) *Coupon {
	usecase := interactor.NewCoupon(
		repository.NewCoupon(),
	)
	return &Coupon{usecase: usecase}
}

// GET /coupons/{id}
func (h *Coupon) Get(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		response.BadRequestErr(w, fmt.Errorf("id param is empty"))
		return
	}

	in := &usecase.CouponGetInput{}

	out, err := h.usecase.Get(r.Context(), in)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	view := newCouponDefaultResponse(out.Coupon)
	response.New(w, view)
}

// POST /coupons
func (h *Coupon) Create(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		response.BadRequestErr(w, fmt.Errorf("error: content length is 0"))
		return
	}

	var j couponCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		response.BadRequestErr(w, err)
		return
	}

	in := &usecase.CouponCreateInput{}

	out, err := h.usecase.Create(r.Context(), in)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	view := newCouponDefaultResponse(out.Coupon)
	response.New(w, view)
}

// GET /coupon_statuses/list
func (h *Coupon) List(w http.ResponseWriter, r *http.Request) {
	in := &usecase.CouponListInput{}

	out, err := h.usecase.List(r.Context(), in)
	if err != nil {
		response.BadRequestErr(w, err)
		return
	}

	view := newCouponListResponse(out.Coupons)
	response.New(w, view)
}

type couponDefaultResponse struct {
	ID             string               `json:"id"`
	Name           string               `json:"name"`
	ExpirationDate uint                 `json:"expiration_date"`
	DiscountRate   uint                 `json:"discount_rate"`
	CreatedAt      string               `json:"created_at"`
	Spot           *spotDefaultResponse `json:"spot"`
}

type couponCreateRequest struct {
	Name           string `json:"name"`
	ExpirationDate uint   `json:"expiration_date"`
	DiscountRate   uint   `json:"discount_rate"`
	SpotID         string `json:"spot_id"`
}

type couponListResponse struct {
	Coupons []*couponDefaultResponse `json:"coupons"`
}

func newCouponDefaultResponse(coupon *model.Coupon) *couponDefaultResponse {
	return &couponDefaultResponse{
		ID:             coupon.ID.String(),
		Name:           coupon.Name,
		ExpirationDate: coupon.ExpirationDate,
		DiscountRate:   coupon.DiscountRate,
		CreatedAt:      coupon.CreatedAt.String(),
		Spot:           newSpotDefaultResponse(coupon.Spot),
	}
}

func newCouponsDefaultResponse(coupons []*model.Coupon) []*couponDefaultResponse {
	var r []*couponDefaultResponse
	for _, coupon := range coupons {
		r = append(r, newCouponDefaultResponse(coupon))
	}
	return r
}

func newCouponListResponse(coupons []*model.Coupon) *couponListResponse {
	return &couponListResponse{
		Coupons: newCouponsDefaultResponse(coupons),
	}
}
