package contract
import "context"

type IHealthCheckUseCase interface {
	Readyness(ctx context.Context) error
	Liveness(ctx context.Context) error
}