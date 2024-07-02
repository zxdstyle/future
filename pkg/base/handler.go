package base

import (
	"context"
	"future-admin/pkg/server/requests"
	"future-admin/pkg/server/responses"
	"github.com/spf13/cast"
)

type Handler[M any] struct {
	logic *Logic[M]
}

func NewHandler[M any](logic *Logic[M]) *Handler[M] {
	return &Handler[M]{
		logic: logic,
	}
}

func (h *Handler[M]) List(ctx context.Context, req *requests.RequestAble) (responses.Response, error) {
	return h.logic.List(ctx, req)
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
	var m M
	if err := req.BodyParser(&m); err != nil {
		return nil, err
	}
	id := req.Params("id")

	if e := h.logic.Update(ctx, cast.ToUint64(id), &m); e != nil {
		return nil, e
	}

	mo, err := h.logic.Show(ctx, cast.ToUint64(id))
	return responses.Success(mo), err
}

func (h *Handler[M]) Show(ctx context.Context, req *requests.RequestAble) (responses.Response, error) {
	id := req.Params("id")
	m, e := h.logic.Show(ctx, cast.ToUint64(id))
	return responses.Success(m), e
}

func (h *Handler[M]) Destroy(ctx context.Context, req *requests.RequestAble) (responses.Response, error) {
	id, err := req.ParamsInt("id")
	if err != nil {
		return nil, err
	}
	if err := h.logic.Destroy(ctx, uint64(id)); err != nil {
		return nil, err
	}
	return responses.Empty(), nil
}
