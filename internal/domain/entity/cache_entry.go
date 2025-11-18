package entity

import (
	"time"

	"github.com/mikiasgoitom/caching-proxy/internal/domain/valueobject"
)

type CacheEntry struct {
	Key valueobject.CacheKey
	Payload ResponseModel
	ExpiresAt time.Time
	StoredAt time.Time
}