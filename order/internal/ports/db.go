package ports

import (
	"github.com/JFMajer/microservices/order/internal/application/core/domain"
)

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}
