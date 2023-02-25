package gateway

import "github.com/vinigofr/golang_ms_wallet/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(*entity.Client) error
}
