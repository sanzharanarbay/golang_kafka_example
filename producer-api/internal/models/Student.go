package models

type Student struct {
	ID    int64   `json:"id" binding:"required"`
	FIO   string  `json:"fio" binding:"required"`
	Group string  `json:"group" binding:"required"`
	Major string  `json:"major" binding:"required"`
	Gpa   float64 `json:"gpa" binding:"required"`
}
