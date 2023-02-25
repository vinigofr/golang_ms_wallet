package createtransaction

import (
	"github.com/vinigofr/golang_ms_wallet/internal/entity"
	"github.com/vinigofr/golang_ms_wallet/internal/gateway"
)

type CreateTransactionInputDTO struct {
	AccountIDFrom string
	AccountIDTo   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	ID string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
}

func NewCreateTransactionUseCase(
	tg gateway.TransactionGateway,
	ag gateway.AccountGateway,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: tg,
		AccountGateway:     ag,
	}
}

func (useCase *CreateTransactionUseCase) Execute(input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err := useCase.AccountGateway.Find(input.AccountIDFrom)
	if err != nil {
		return nil, err
	}

	accountTo, err := useCase.AccountGateway.Find(input.AccountIDTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}

	if err = useCase.TransactionGateway.Create(transaction); err != nil {
		return nil, err
	}

	return &CreateTransactionOutputDTO{
		ID: transaction.ID,
	}, nil

}
