package createaccount

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vinigofr/golang_ms_wallet/internal/entity"
)

// Client mock
type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)

	return args.Get(0).(*entity.Client), args.Error(1)
}
func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

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

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("Vinicius", "v@v.com")
	clientMock := &ClientGatewayMock{}
	clientMock.On("Get", client.ID).Return(client, nil)

	accountMock := &AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	useCase := NewCreateAccountUseCase(accountMock, clientMock)
	inputDto := CreateAccountInputDTO{
		ClientID: client.ID,
	}

	output, err := useCase.Execute(inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output.ID)

	clientMock.AssertExpectations(t)
	accountMock.AssertExpectations(t)

	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)

}
