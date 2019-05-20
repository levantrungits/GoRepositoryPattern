package repository

import (
	models "go-repo-modules/model"
)

// ~ user_repo.go
type UserRepo interface {
	Select() ([]models.User, error)
	Insert(u models.User) (error)
}