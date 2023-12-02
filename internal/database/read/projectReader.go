package read

import (
	"backendAleshinaFeklistova/internal/models"
)

func (pr *ProjectReadRepository) GetAll() []models.Project {
	res := TestDataProj
	return res
}

func (pr *ProjectReadRepository) GetById(id int) models.Project {
	res, _ := executeSqlProject(id, "")
	return res[0]
}

func (pr *ProjectReadRepository) GetProjectsByType(prType string) []models.Project {
	res, _ := executeSqlProject(0, prType)
	return res
}

func executeSqlProject(id int, prType string) ([]models.Project, error) {
	var res []models.Project
	if id == 0 {
		for _, project := range TestDataProj {
			if project.ProjectType == prType {
				res = append(res, project)
			}
		}
	} else {
		for _, project := range TestDataProj {
			if project.ProjectId == id {
				res = append(res, project)
				return res, nil
			}
		}
	}

	return res, nil
}

var TestDataProj = []models.Project{
	{
		ProjectId:   1,
		ProjectName: "test1",
		ProjectType: "Логистическая регрессия",
		Target:      "s1",
		Status:      "Error",
		ResultLink:  "",
	},
	{
		ProjectId:   2,
		ProjectName: "test2",
		ProjectType: "Метод опорных векторов",
		Target:      "s2",
		Status:      "In progress",
		ResultLink:  "",
	},
	{
		ProjectId:   3,
		ProjectName: "test3",
		ProjectType: "Линейная регрессия",
		Target:      "s3",
		Status:      "Finished",
		ResultLink:  "link.link",
	}}
