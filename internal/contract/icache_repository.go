package contract

import (
	"context"

	"github.com/mikiasgoitom/caching-proxy/internal/domain/entity"
	"github.com/mikiasgoitom/caching-proxy/internal/domain/valueobject"
)
type ICacheRepository interface {
	Get(ctx context.Context, key valueobject.CacheKey) (entity.CacheEntry, bool, error)
	Set(ctx context.Context, key valueobject.CacheKey, value entity.CacheEntry) error
	Delete(ctx context.Context, key valueobject.CacheKey) error
	HealthCheck(ctx context.Context) error
}