package models

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/google/uuid"
)

func NewUserDomain(
	email string,
	name string,
	password string,
) UserDomainInterface {
	return &userDomain{
		id:       uuid.New(),
		email:    email,
		name:     name,
		password: password,
	}
}

type UserDomainInterface interface {
	GetId() uuid.UUID
	GetEmail() string
	GetName() string
	GetPassword() string
	EncryptPassword()
}

type userDomain struct {
	id       uuid.UUID
	email    string
	name     string
	password string
}

func (ud *userDomain) EncryptPassword() {
	encrypt := md5.New()
	defer encrypt.Reset()
	encrypt.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(encrypt.Sum(nil))
}

func (ud *userDomain) GetId() uuid.UUID {
	return ud.id
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}
