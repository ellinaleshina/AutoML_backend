package models

type Review struct {
	Precision string `json:"precision"`
	Recall    string `json:"recall"`
	F1        string `json:"f_1"`
}
