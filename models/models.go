package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	autoId      bson.ObjectId `json:"auto_id,omitempty" bson:"auto_id,omitempty"`
	id          string        `json:"id" form:"id" binding:"required" bson:"id"`
	password    string        `json:"password form:"password" binding:"required" bson:"password"`
	name        string        `json:"name" form:"name" binding:"required" bson:"name"`
	studentNum  int16         `json:"student_num" form:"student_num" binding:"required" bson:"student_num"`
}