package contract
import "context"

type IHealthCheckUseCase interface {
	Ready(ctx context.Context) error
	Live(ctx context.Context) error
}