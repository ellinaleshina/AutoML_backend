package read

type Reader interface {
	GetAll() any
	GetById(id int) any
}

type UserReadRepository struct {
	sqlQuery string
}

type ProjectReadRepository struct {
	sqlQuery string
}

func NewUserReadRepository() *UserReadRepository {
	return &UserReadRepository{
		sqlQuery: ""}
}

func NewProjectReadRepository() *ProjectReadRepository {
	return &ProjectReadRepository{
		sqlQuery: ""}
}
