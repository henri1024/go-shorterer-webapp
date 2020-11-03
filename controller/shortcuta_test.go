package controller

import (
	"shorterer/model"

	"github.com/stretchr/testify/mock"
)

type MocKRepository struct {
	mock.Mock
}

func (m *MocKRepository) Save(model *model.ShortLink, flag bool) error {
	args := m.Called(model, flag)

	return args.Error(0)
}

func (m *MocKRepository) Delete(string) error {
	return nil
}

func (m *MocKRepository) GetDestination(key string) (string, error) {
	args := m.Called(key)

	val := args.Get(0).(string)
	err := args.Error(1)
	return val, err
}
