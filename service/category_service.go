package service

import (
  "go-unittest/repository"
  "go-unittest/entity"
  "errors"
)
// idealnya bikin kontrak interface jg, cmn utk mempermudah latian pake struct lgsg
type CategoryService struct {
  Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
  category := service.service.Repository.FindById(id)
  if category == nil {
    return category, errors.New("Category not found")
  } else {
    return category, nil
  }
}