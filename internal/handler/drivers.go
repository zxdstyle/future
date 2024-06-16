package handler

import (
	"context"
	"future-admin/internal/logic"
	"future-admin/internal/logic/driver"
	"future-admin/pkg/server/requests"
	"future-admin/pkg/server/responses"
)

var Driver = &driverHandler{}

type driverHandler struct {
}

func (h *driverHandler) List(ctx context.Context, req *requests.RequestAble) (responses.Response, error) {
	drivers, err := logic.Invoke[*driver.Logic]().List(ctx)
	if err != nil {
		return nil, err
	}
	return responses.Success(drivers), nil
}
