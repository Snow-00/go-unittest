package service

import (
	"testing"
	"go-unittest/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

// testing dg harapan hasil nya nil
func TestCategoryService_GetNotFound(t *testing.T) {
	// program mock
	// ketika func FindById (func di repository) dipanggil oleh mock dg param 1 maka return datanya nil
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

// testing dg harapan hasilnya ada / datanya sukses muncul
func TestCategoryService_GetSuccess(t *testing.T) {
  category := entity.Category{
    Id:   "1",
    Name: "Laptop",
  }
  categoryRepository.Mock.On("FindById", "2").Return(category)

  result, err := categoryService.Get("2")
  assert.Nil(t, err)
  assert.NotNil(t, result)
  assert.Equal(t, category.Id, result.Id)
  assert.Equal(t, category.Name, result.Name)
}