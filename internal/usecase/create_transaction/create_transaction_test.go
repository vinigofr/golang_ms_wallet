package createtransaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vinigofr/golang_ms_wallet/internal/entity"
)

// Account mock
type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}
func (m *AccountGatewayMock) Find(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

// Transaction mock
type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func TestCreateTransaction_Execute(t *testing.T) {
	client1, _ := entity.NewClient("Vinicius", "v@v.com")
	client2, _ := entity.NewClient("Gouveia", "g@g.com")
	account1 := entity.NewAccount(client1)
	account2 := entity.NewAccount(client2)
	account1.Credit(1000.00)
	account2.Credit(1000.00)

	accountGatewayMock := &AccountGatewayMock{}
	accountGatewayMock.On("Find", mock.AnythingOfType("string")).Return(account1, nil)
	accountGatewayMock.On("Find", mock.AnythingOfType("string")).Return(account2, nil)

	transactionGatewayMock := &TransactionGatewayMock{}
	transactionGatewayMock.On("Create", mock.Anything).Return(nil)

	useCase := NewCreateTransactionUseCase(transactionGatewayMock, accountGatewayMock)
	inputDto := &CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        1000.00,
	}

	outputDto, err := useCase.Execute(*inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, outputDto)

	accountGatewayMock.AssertExpectations(t)
	transactionGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertNumberOfCalls(t, "Find", 2)
	transactionGatewayMock.AssertNumberOfCalls(t, "Create", 1)
}
