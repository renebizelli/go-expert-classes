package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {

	name := "René Bizelli"
	email := "rene@rene.com"
	password := "123456"

	user, e := NewUser(name, email, password)

	assert.Nil(t, e)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Id)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)
}

func TestUserValidateTes(t *testing.T) {
	name := "René Bizelli"
	email := "rene@rene.com"
	password := "123456"

	user, e := NewUser(name, email, password)

	assert.Nil(t, e)

	assert.True(t, user.ValidatePassword(password))
	assert.False(t, user.ValidatePassword(password+"1"))
	assert.NotEqual(t, password, user.Password)

}
