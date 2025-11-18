package entity

import (
	"net/http"
	"net/url"
	"time"

	"github.com/mikiasgoitom/caching-proxy/internal/domain/valueobject"
)

type RequestModel struct {
	ID         string
	Method     string
	URL        *url.URL
	Headers    valueobject.HeaderSet
	Body       []byte
	ReceivedAt time.Time
}