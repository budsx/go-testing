package service

import (
	"errors"
	"go-testing/entity"
	"go-testing/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("Catogory not found")
	} else {
		return category, nil
	}

}
