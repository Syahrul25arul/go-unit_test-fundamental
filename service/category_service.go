package service

import (
	"belajar-golan-unit-test/entity"
	"belajar-golan-unit-test/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category Not Found")
	} else {
		return category, nil
	}
}
