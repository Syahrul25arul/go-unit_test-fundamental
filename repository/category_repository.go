package repository

import "belajar-golan-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
