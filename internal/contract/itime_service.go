package contract

import "time"

type ITimeService interface {
	Now() time.Time
	NowUnix() int64
}
