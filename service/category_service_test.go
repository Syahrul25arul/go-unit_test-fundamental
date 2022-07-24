package service

import (
	"belajar-golan-unit-test/entity"
	"belajar-golan-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

var Assert *assert.Assertions = assert.New(&testing.T{})

func TestCategoryServiceNotFound(t *testing.T) {
	// program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	Assert.Nil(category)
	Assert.NotNil(err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "Laptop",
	}
	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.Get("2")
	Assert.Nil(err)
	Assert.NotNil(result)
	Assert.Equal(category.Id, result.Id)
	Assert.Equal(category.Name, result.Name)
}
