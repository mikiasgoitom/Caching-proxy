package entity

import (
	"net/http"
	"time"
)

type ResponseModel struct {
	ID     string
	Status int
	Header http.Header
	Body   []byte
	GeneratedAt time.Time
	Cacheablee bool
}