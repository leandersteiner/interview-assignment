package calculator

import (
	"context"
	"github.com/leandersteiner/interview-assignment/internal/web"
	"net/http"
)

type getter interface {
	Get(p Pagination) PaginatedResult[[]Result]
}

type Handler struct {
	service *Service
	getter  getter
}

func NewHandler(service *Service, getter getter) *Handler {
	return &Handler{
		service,
		getter,
	}
}

func (h *Handler) Addition(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	req := &AdditionRequest{}
	if err := web.Decode(r, req); err != nil {
		return web.NewError(http.StatusBadRequest, err.Error())
	}

	result, err := h.service.Add(req.SummandOne, req.SummandTwo)
	if err != nil {
		return web.NewError(http.StatusBadRequest, err.Error())
	}

	resp := &AdditionResponse{
		Sum: result.Value,
	}
	return web.Respond(ctx, w, resp, http.StatusOK)
}

func (h *Handler) Subtraction(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	req := &SubtractionRequest{}
	if err := web.Decode(r, req); err != nil {
		return web.NewError(http.StatusBadRequest, err.Error())
	}

	result, err := h.service.Sub(req.Minuend, req.Subtrahend)
	if err != nil {
		return web.NewError(http.StatusBadRequest, err.Error())
	}

	resp := &SubtractionResponse{
		Difference: result.Value,
	}
	return web.Respond(ctx, w, resp, http.StatusOK)
}

func (h *Handler) Multiplication(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	req := &MultiplicationRequest{}
	if err := web.Decode(r, req); err != nil {
		return web.NewError(http.StatusBadRequest, err.Error())
	}

	result, err := h.service.Mul(req.FactorOne, req.FactorTwo)
	if err != nil {
		return web.NewError(http.StatusBadRequest, err.Error())
	}

	resp := &MultiplicationResponse{
		Product: result.Value,
	}
	return web.Respond(ctx, w, resp, http.StatusOK)
}

func (h *Handler) Division(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	req := &DivisionRequest{}
	if err := web.Decode(r, req); err != nil {
		return web.NewError(http.StatusBadRequest, err.Error())
	}

	result, err := h.service.Div(req.Dividend, req.Divisor)
	if err != nil {
		return web.NewError(http.StatusBadRequest, err.Error())
	}

	resp := &DivisionResponse{
		Quotient: result.Value,
	}
	return web.Respond(ctx, w, resp, http.StatusOK)
}

func (h *Handler) GetRecent(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	pagination := parsePagination(r)

	results := h.getter.Get(pagination)

	expressions := make([]string, len(results.Result))
	for i, result := range results.Result {
		expressions[i] = result.Expression
	}

	resp := RecentResponse{
		Results:  expressions,
		Metadata: results.Metadata,
	}

	return web.Respond(ctx, w, resp, http.StatusOK)
}
