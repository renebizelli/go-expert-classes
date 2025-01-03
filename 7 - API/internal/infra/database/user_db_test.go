package database

import (
	"renebizelli/api/internal/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {

	db, e := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if e != nil {
		t.Error(e)
	}

	db.AutoMigrate(&entities.User{})

	name := "René Bizelli"
	email := "rene@rene.com"
	password := "123456"

	user, _ := entities.NewUser(name, email, password)
	userDb := NewUser(db)

	eUserDb := userDb.Create(user)
	assert.Nil(t, eUserDb)

	var userFound entities.User

	err := db.First(&userFound, " id = ?", user.Id).Error
	assert.Nil(t, err)

	assert.Equal(t, name, userFound.Name)
	assert.Equal(t, email, userFound.Email)
	assert.NotEmpty(t, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db, e := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if e != nil {
		t.Error(e)
	}

	db.AutoMigrate(&entities.User{})

	name := "René Bizelli"
	email := "rene@rene.com"
	password := "123456"

	user, _ := entities.NewUser(name, email, password)
	userDb := NewUser(db)

	eUserDb := userDb.Create(user)
	assert.Nil(t, eUserDb)

	userFound, err := userDb.FindByEmail(email)
	assert.Nil(t, err)

	assert.Equal(t, name, userFound.Name)
	assert.Equal(t, email, userFound.Email)
	assert.NotEmpty(t, userFound.Password)
}
