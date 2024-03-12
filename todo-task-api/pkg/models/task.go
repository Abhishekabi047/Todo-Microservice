package models

type Tasks struct{
	Id int64 `json:"id" gorm:"primaryKey"`
	Task string `json:"task"`
	Description string `json:"description"`
	Done bool `json:"done" gorm:"default:false"`
}