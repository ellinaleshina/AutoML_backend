package service

import (
	"backendAleshinaFeklistova/internal/database/read"
	"backendAleshinaFeklistova/internal/database/write"
	"backendAleshinaFeklistova/internal/models"
	"io"
	"net/http"
	"os"
)

var (
	pr = read.NewProjectReadRepository()
	pw = write.NewProjectWriteRepository()
)

func CreateProject(r *http.Request) error {
	var proj models.Project
	reader, _ := r.MultipartReader()
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if part.FileName() != "" {
			file, err := os.Create(part.FileName())
			if err != nil {
				return err
			}
			defer file.Close()
			// Копируем содержимое загруженного файла в новый файл на сервере
			_, err = io.Copy(file, part)
			if err != nil {
				return err
			}
		} else {
			formName := part.FormName()
			formValue, err := io.ReadAll(part)
			if err != nil {
				return err
			}

			switch formName {
			case "ProjectName":
				proj.ProjectName = string(formValue)
			case "ProjectType":
				proj.ProjectType = string(formValue)
			case "ProjectTarget":
				proj.Target = string(formValue)

			}
		}
	}
	pw.Project = proj
	pw.Create()
	//proj.ProjectId = read.TestDataProj[len(read.TestDataProj)-1].ProjectId + 1
	//read.TestDataProj = append(read.TestDataProj, proj)
	return nil
}

func GetProjectById(id int) models.Project {
	return pr.GetById(id)
}

func GetAllProjects() []models.Project {
	return pr.GetAll()
}

func EditProject(r *http.Request) error {
	var proj models.Project
	reader, _ := r.MultipartReader()
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if part.FileName() != "" {
			file, err := os.Create(part.FileName())
			if err != nil {
				return err
			}
			defer file.Close()
			// Копируем содержимое загруженного файла в новый файл на сервере
			_, err = io.Copy(file, part)
			if err != nil {
				return err
			}
		} else {
			formName := part.FormName()
			formValue, err := io.ReadAll(part)
			if err != nil {
				return err
			}

			switch formName {
			case "ProjectName":
				proj.ProjectName = string(formValue)
			case "ProjectType":
				proj.ProjectType = string(formValue)
			case "ProjectTarget":
				proj.Target = string(formValue)

			}
		}
	}
	pw.Project = proj
	pw.Edit()
	return nil
}

func DeleteProject(id int) {
	pw.ProjectId = id
	pw.Delete()
}

func GetProjectsByType(prType string) []models.Project {
	return pr.GetProjectsByType(prType)
}
