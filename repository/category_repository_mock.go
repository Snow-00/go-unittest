package repository

import (
  "github.com/stretchr/testify/mock"
  "go-unittest/entity"
)

type CategoryRepositoryMock struct {
  Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
  arguments := repository.Mock.Called(id)  // Called(param func FindById)
}