package models

import (
"gopkg.in/mgo.v2/bson"
)

type (
	// User represents the structure of our resource
	User struct {
		Id     bson.ObjectId `json:"id" bson:"_id"`
		Session     string `json:"session" bson:"session"`
		Email  string        `json:"email" bson:"email"`
		Name   string        `json:"name" bson:"name"`
		Password    string           `json:"password" bson:"password"`
		Projects    []bson.ObjectId           `json:"projects" bson:"projects"`
		Success bool `json:"success" bson:"success"`
		LoggedIn bool `json:"loggedin" bson:"loggedin"`
	}

	UserSession struct {
		Id     bson.ObjectId `json:"id" bson:"_id"`
		Session     string `json:"session" bson:"session"`
		Success bool `json:"success"`
	}

	Project struct {
		Id     bson.ObjectId `json:"id" bson:"_id"`
		Session     string `json:"session" bson:"session"`
		Name   string        `json:"name" bson:"name"`
		Client   string        `json:"client" bson:"client"`
		Entries    []bson.ObjectId           `json:"entires" bson:"entries"`
	}

	Entry struct {
		Id     bson.ObjectId `json:"id" bson:"_id"`
		Session     string `json:"session" bson:"session"`
		Status   string        `json:"status" bson:"status"`
		Rate   int        `json:"rate" bson:"rate"`
		StartYear   int        `json:"startYear" bson:"startYear"`
		StartMonth   int        `json:"startMonth" bson:"startMonth"`
		StartDay   int        `json:"startDay" bson:"startDay"`
		StartHour   int        `json:"startHour" bson:"startHour"`
		StartMin   int        `json:"startMin" bson:"startMin"`
		EndYear   int        `json:"endYear" bson:"endYear"`
		EndMonth   int        `json:"endMonth" bson:"endMonth"`
		EndDay   int        `json:"endDay" bson:"endDay"`
		EndHour   int        `json:"endHour" bson:"endHour"`
		EndMin   int        `json:"endMin" bson:"endMin"`
		Category   string        `json:"category" bson:"category"`
		Subcategory   string        `json:"Subcategory" bson:"Subcategory"`
		Info   string        `json:"name" bson:"name"`
	}
)
