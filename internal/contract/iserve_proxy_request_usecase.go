package contract

import (
	"context"

	"github.com/mikiasgoitom/caching-proxy/internal/domain/entity"
)

type IServeProxyRequestUseCase interface {
	ServeProxyRequest(ctx context.Context, req entity.RequestModel) (entity.ResponseModel,error)
}