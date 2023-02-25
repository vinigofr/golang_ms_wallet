package gateway

import "github.com/vinigofr/golang_ms_wallet/internal/entity"

type AccountGateway interface {
	Save(account entity.Account) error
	Find(id string) (*entity.Account, error)
}
