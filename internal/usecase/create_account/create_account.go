package createaccount

import (
	"github.com/vinigofr/golang_ms_wallet/internal/entity"
	"github.com/vinigofr/golang_ms_wallet/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientID string
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClienteGateway gateway.ClientGateway
}

func NewCreateAccountUseCase(
	account gateway.AccountGateway,
	client gateway.ClientGateway,
) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: account,
		ClienteGateway: client,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := uc.ClienteGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}

	account := *entity.NewAccount(client)

	err = uc.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}

	return &CreateAccountOutputDTO{
		ID: account.ID,
	}, nil
}
