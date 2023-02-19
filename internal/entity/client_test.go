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
