package gateway

import "github.com/vinigofr/golang_ms_wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
