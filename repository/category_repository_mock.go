package repository

import (
	"github.com/raihanmd/app-golang-module/entity"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repo *CategoryRepositoryMock) FindById(id string) *entity.Category{
	arguments := repo.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	}else {
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}