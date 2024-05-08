package base

import (
	"context"
	"future-admin/pkg/server/requests"
	"future-admin/pkg/server/responses"
)

type Handler[M any] struct {
	logic *Logic[M]
}

func NewHandler[M any](logic *Logic[M]) *Handler[M] {
	return &Handler[M]{
		logic: logic,
	}
}

func (h *Handler[M]) Create(ctx context.Context, req *requests.RequestAble) (responses.Response, error) {
	var m M

	if err := req.Validate(&m); err != nil {
		return nil, err
	}

	if err := h.logic.Create(ctx, &m); err != nil {
		return responses.Internal(err.Error()), nil
	}
	return responses.Success(m), nil
}

func (h *Handler[M]) Update(ctx context.Context, req *requests.RequestAble) (responses.Response, error) {
	return nil, nil
}

func (h *Handler[M]) Show(ctx context.Context, req *requests.RequestAble) (responses.Response, error) {
	return nil, nil
}

func (h *Handler[M]) Destroy(ctx context.Context, req *requests.RequestAble) (responses.Response, error) {
	return nil, nil
}
