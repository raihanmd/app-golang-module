package repository

import "github.com/raihanmd/app-golang-module/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}