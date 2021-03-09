package models

import (
	"time"
)

type User struct {
	Id          int        `json:"id"`
	UserId      string     `json:"userid"`
	Password    string     `json:"password"`
	Name        string     `json:"name"`
	StudentNum  int        `json:"student_num"`
}

type Board struct{
	Id         int         `json:"id"`
	Title      string      `json:"title"`
	Content    string      `json:"content"`
	Author     string      `json:"author"`
	AuthorNum  int         `json:"author_num"'`
	Date       time.Time   `json:"date"`
	PageView   int         `json:"page_view"`
}