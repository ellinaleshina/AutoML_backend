package read

import "backendAleshinaFeklistova/internal/models"

func (ur *UserReadRepository) GetAll() []models.User {
	res := TestData
	return res
}

func (ur *UserReadRepository) GetById(id int) models.User {
	res, _ := executeSqlUser(id, "")
	return res
}

func (ur *UserReadRepository) ExistsByEmail(email string) bool {
	_, res := executeSqlUser(0, email)
	return res
}

func (ur *UserReadRepository) GetByEmail(email string) models.User {
	res, _ := executeSqlUser(0, email)
	return res
}

func (ur *UserReadRepository) GetHashByEmail(email string) string {
	res, _ := executeSqlUser(0, email)
	return res.Password
}

func executeSqlUser(id int, email string) (models.User, bool) {
	if email == "" {
		for _, datum := range TestData {
			if datum.Id == id {
				return datum, false
			}
		}
	} else {
		for _, datum := range TestData {
			if datum.Email == email {
				return datum, true
			}
		}
	}
	return models.User{}, false
}

var TestData = []models.User{
	{
		Id:       1,
		Email:    "test1",
		Password: "$2a$10$df0LnDx0QNcKaDy0M29tI.om9cjWUqAmqY.nB3RQ2jujReoOSFLfS",
		Name:     "Name Surname",
		Phone:    "89123456789",
		Projects: nil,
	},
	{
		Id:       2,
		Email:    "ivanivan@mail.ru",
		Password: "$2a$10$df0LnDx0QNcKaDy0M29tI.om9cjWUqAmqY.nB3RQ2jujReoOSFLfS",
		Name:     "Иван Иванов",
		Phone:    "89098765432",
		Projects: nil,
	},
	{
		Id:       3,
		Email:    "petrpetr@mail.ru",
		Password: "$2a$10$df0LnDx0QNcKaDy0M29tI.om9cjWUqAmqY.nB3RQ2jujReoOSFLfS",
		Name:     "Петр Петров",
		Phone:    "89433154367",
		Projects: nil,
	},
}
