package model

type Todo struct {
	ModelBase
	Title     string `gorm:"" json:"title"`
	Completed bool   `gorm:"" json:"completed"`
}
