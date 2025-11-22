package contract

import "time"

type ICacheControl interface {
	Parse(header string) map[string]string
	Has(directive string) bool
	Get(directive string) (string, bool)
	GetDuration(directive string) (time.Duration, bool)
}