package service

import (
	"backendAleshinaFeklistova/internal/database/read"
	"backendAleshinaFeklistova/internal/database/write"
	"backendAleshinaFeklistova/internal/models"
)

var (
	ur = read.NewUserReadRepository()
	uw = write.NewUserWriteRepository()
)

func RegisterUser(user models.User) {
	uw.User = user
	uw.Create()
}

func ExistsByEmail(email string) bool {
	return ur.ExistsByEmail(email)
}

func GetUserById(id int) models.User {
	return ur.GetById(id)
}

func EditProfile(user models.User) {
	uw.User = user
	uw.Edit()
}

func DeleteProfile(id int) {
	uw.Id = id
	uw.Delete()
}

func GetAllUsers() []models.User {
	return ur.GetAll()
}

func getPasswordHash(email string) string {
	return ur.GetHashByEmail(email)
}
