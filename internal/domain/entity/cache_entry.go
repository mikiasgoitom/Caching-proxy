package entity

import (
	"github.com/mikiasgoitom/caching-proxy/internal/domain/valueobject"
)

type CacheEntry struct {
	Key       valueobject.CacheKey
	Payload   ResponseModel
	ExpiresAt int64
	StoredAt  int64
}
