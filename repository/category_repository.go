package repository

import "go-testing/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
