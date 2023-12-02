package write

import "backendAleshinaFeklistova/internal/models"

type Writer interface {
	Create()
	Edit()
	Delete()
}

type UserWriteRepository struct {
	sqlQuery string
	models.User
}

type ProjectWriteRepository struct {
	sqlQuery string
	models.Project
}

func NewUserWriteRepository() *UserWriteRepository {
	return &UserWriteRepository{
		sqlQuery: "",
		User:     models.User{},
	}
}

func NewProjectWriteRepository() *ProjectWriteRepository {
	return &ProjectWriteRepository{
		sqlQuery: "",
		Project:  models.Project{},
	}
}
