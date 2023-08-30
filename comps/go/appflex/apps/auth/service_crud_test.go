package auth

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockedDao struct {
	DaoDatabase
	users map[string]Model
}

func newMockedDao() *mockedDao {
	r := &mockedDao{}
	r.users = make(map[string]Model)
	return r
}

func (d *mockedDao) Create(v Model) error {
	d.users[v.Email] = v
	return nil
}

func (d mockedDao) Update(v Model) error {
	return d.DaoDatabase.Update(v)
}

func (d mockedDao) Delete(v Model) error {
	return d.DaoDatabase.Delete(v)
}

func (d *mockedDao) Find(v Query) (r []Model, err error) {
	if v.Email != "" {
		if d.users[v.Email].Id != "" {
		}
	}
	return
}

func (d *mockedDao) FindByEmail(email string) (r Model, err error) {
	return
}

func TestRegisterValid(t *testing.T) {
	service, err := NewService()
	service.dao = newMockedDao()
	assert.Nil(t, err)
	testCases := []struct {
		name, email, password string
		age                   uint
		expected              string
	}{
		{name: "Jonh", email: "jonh@test.com", password: "12345678", age: 15, expected: ""},
		{name: "Jonh", email: "jonh@test.com", password: "12345678", age: 10, expected: ""},
		{name: "Jonh", email: "jonh@test.com", password: "12345678", age: 100, expected: ""},
	}
	for _, testCase := range testCases {
		testname := fmt.Sprintf("name: '%s', email: '%s', expected: '%s'", testCase.name, testCase.email, testCase.expected)
		t.Run(testname, func(t *testing.T) {
			err = service.register(testCase.name, testCase.email, testCase.password, testCase.age)
			assert.Nil(t, err)
		})
	}
}

func TestRegisterSuccess(t *testing.T) {
	service, err := NewService()
	service.dao = newMockedDao()
	assert.Nil(t, err)
	err = service.register("jonh", "jonh@test.com", "123456", 15)
	assert.Nil(t, err)
}

func TestRegisterInvalid(t *testing.T) {
	service, err := NewService()
	service.dao = &mockedDao{}
	assert.Nil(t, err)
	testCases := []struct {
		name, email, password string
		age                   uint
		expected              string
	}{
		{name: "a", email: "jonh@test.com", password: "123456", expected: "Name (A) precisa ter entre 2 caracteres e 120, ele possui 1,Age precisa ser entre entre 10 e 100, valor atual 0"},
		{name: "maria", email: "maria", password: "123456", age: 100, expected: "Campo Email \"maria\" nao é um email válido,Email (maria) precisa ter entre 8 caracteres e 120, ele possui 5"},
		{name: "Jonh", email: "jonh@test.com", password: "1", age: 10, expected: "Senha (1) precisa ter entre 4 caracteres e 30, ele possui 1"},
		{name: "Jonh", email: "jonh@test.com", password: "12345678", age: 9, expected: "Age precisa ser entre entre 10 e 100, valor atual 9"},
	}
	for _, testCase := range testCases {
		testname := fmt.Sprintf("name: '%s', email: '%s', expected: '%s'", testCase.name, testCase.email, testCase.expected)
		t.Run(testname, func(t *testing.T) {
			err = service.register(testCase.name, testCase.email, testCase.password, testCase.age)
			assert.Equal(t, testCase.expected, err.Error())
		})
	}
}
