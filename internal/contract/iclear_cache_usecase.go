package contract

import "context"

type IClearCacheUseCase interface {
	ClearCache(ctx context.Context) error
}