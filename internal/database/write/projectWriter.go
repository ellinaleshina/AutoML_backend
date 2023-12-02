package write

import (
	"backendAleshinaFeklistova/internal/database/read"
	"backendAleshinaFeklistova/internal/models"
)

func (pw *ProjectWriteRepository) Create() {
	executeSqlProject(1, pw)
}

func (pw *ProjectWriteRepository) Edit() {
	executeSqlProject(2, pw)
}

func (pw *ProjectWriteRepository) Delete() {
	executeSqlProject(3, pw)
}

func executeSqlProject(operId int, pw *ProjectWriteRepository) {
	if operId == 1 {
		read.TestDataProj = append(read.TestDataProj, pw.Project)
	}
	if operId == 2 {
		for i, datum := range read.TestDataProj {
			if datum.ProjectId == pw.Project.ProjectId {
				read.TestDataProj[i] = pw.Project
			}
		}
	}
	if operId == 3 {
		for i, datum := range read.TestDataProj {
			if datum.ProjectId == pw.ProjectId {
				read.TestDataProj[i] = read.TestDataProj[len(read.TestDataProj)-1]
				read.TestDataProj[len(read.TestDataProj)-1] = models.Project{}
				read.TestDataProj = read.TestDataProj[:len(read.TestDataProj)-1]
			}
		}
	}
}
