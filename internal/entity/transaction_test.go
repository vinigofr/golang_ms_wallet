package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("João", "j@j.com")
	account1 := NewAccount(client1)
	account1.Credit(100.00)

	client2, _ := NewClient("Maria", "m@m.com")
	account2 := NewAccount(client2)
	account2.Credit(100.00)

	transaction, err := NewTransaction(account1, account2, 50.00)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)

	assert.Equal(t, 50.00, account1.Balance)
	assert.Equal(t, 150.00, account2.Balance)
}

func TestCreateTransactionWithNoBalance(t *testing.T) {
	client1, _ := NewClient("João", "j@j.com")
	account1 := NewAccount(client1)
	account1.Credit(100.00)

	client2, _ := NewClient("Maria", "m@m.com")
	account2 := NewAccount(client2)
	account2.Credit(100.00)

	transaction, err := NewTransaction(account1, account2, 2000.00)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)

	assert.Equal(t, "insufficient founds", err.Error())
}
