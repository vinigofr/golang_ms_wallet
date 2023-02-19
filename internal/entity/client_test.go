package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Vinicius Gouveia", "j@j.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Vinicius Gouveia", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestCreateNewClientWithInvalidArgs(t *testing.T) {
	client, err := NewClient("", "")

	assert.Nil(t, client)
	assert.NotNil(t, err)
	assert.Equal(t, "name is requird", client.Name)
	assert.Equal(t, "email is required", client.Email)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Vinicius Gouveia", "j@j.com")

	err := client.Update("Vinicius Updated", "a@a.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Vinicius Updated", client.Name)
	assert.Equal(t, "a@a.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("Vinicius Gouveia", "j@j.com")
	assert.NotNil(t, client)

	err := client.Update("", "a@a.com")
	assert.Equal(t, "name is required", err.Error())

	err = client.Update("Name", "")
	assert.Equal(t, "email is required", err.Error())
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("Vinicius Gouveia", "a@a.com")
	account := NewAccount(client)

	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
