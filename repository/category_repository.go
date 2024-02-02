package repository

import "unit-test-go-pzn/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
