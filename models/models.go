package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id          bson.ObjectId  `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId      string         `json:"user_id,omitempty" form:"user_id" binding:"required" bson:"user_id,omitempty"`
	Password    string         `json:"password form:"password" binding:"required" bson:"password"`
	Name        string         `json:"name" form:"name" binding:"required" bson:"name"`
	StudentNum  int32          `json:"student_num" form:"student_num" binding:"required" bson:"student_num"`
}

type Board struct{
	Id         bson.ObjectId  `json:"_id,omitempty" bson:"_id,omitempty"`
	Title      string         `json:"title" form:"title" binding:"required" bson:"title"`
	Content    string         `json:"content" form:"content" bson:"content"`
	Author     string         `json:"author" form:"author" binding"required" bson"author"`
	AuthorNum  int32          `json:"author_num" form:"author_num" binding:"required" bson:"author_num"'`
	Date       time.Time      `json:"date" form:"date" bson:"date"`
	PageView   int64          `json:"page_view" form:"page_view" bson:"page_view"`
}

func (s *Board) HandlePageView() {
	s.PageView += 1
}