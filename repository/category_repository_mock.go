package repository

import (
	"go-testing/entity"

	"github.com/stretchr/testify/mock"
)

type CatogoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CatogoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.Category)
		return &category

	}
}
