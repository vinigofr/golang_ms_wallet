package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("Vinicius Gouveia", "a@a.com")
	account := NewAccount(client)

	assert.NotNil(t, account)
	assert.Equal(t, account.ID, client.ID)
}

func TestCreateAccountWithInvalidClient(t *testing.T) {
	err := NewAccount(nil)
	assert.Nil(t, err)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("Vinicius Gouveia", "a@a.com")

	account := NewAccount(client)
	account.Credit(100.00)

	assert.Equal(t, 100.00, account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("Vinicius Gouveia", "a@a.com")

	account := NewAccount(client)
	account.Credit(500.00)
	account.Debit(250.00)

	assert.Equal(t, 250.00, account.Balance)
}
