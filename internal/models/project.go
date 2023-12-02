package models

type Project struct {
	ProjectId   int    `json:"project_id" swaggerignore:"true"`
	ProjectName string `json:"project_name"`
	ProjectType string `json:"project_type"`
	Target      string `json:"target"`
	Status      string `json:"status"`
	ResultLink  string `json:"result_link"`
}
