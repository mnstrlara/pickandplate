package model

type Tags struct {
	Id   int    `gorm:"type:int;primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}
