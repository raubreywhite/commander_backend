package models

import (
"gopkg.in/mgo.v2/bson"
)

type (

	UserSession struct {
		Id     bson.ObjectId `json:"id" bson:"_id"`
		Session     string `json:"session" bson:"session"`
		Success bool `json:"success"`
	}

	Entry struct {
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

  Bill struct {
		Year int `json:"year" bson:"year"`
		Month int `json:"month" bson:"month"`
		Billed string `json:"billed" bson:"billed"`
		MoneyBilled int `json:"moneybilled" bson:"moneybilled"`
		MoneyReceived int `json:"moneyreceived" bson:"moneyreceived"`
	}

  Project struct {
		Name   string        `json:"name" bson:"name"`
		Entries    []Entry           `json:"entries" bson:"entries"`
	}

  Client struct {
    Name   string        `json:"name" bson:"name"`
    Bills []Bill `json:"bills" bson:"bills"`
    Projects    []Project           `json:"projects" bson:"projects"`
  }

  // User represents the structure of our resource
	User struct {
		Id     bson.ObjectId `json:"id" bson:"_id"`
		Session     string `json:"session" bson:"session"`
		Email  string        `json:"email" bson:"email"`
		Name   string        `json:"name" bson:"name"`
    Clients    []Client           `json:"clients" bson:"clients"`
		Password    string           `json:"password" bson:"password"`
		Success bool `json:"success" bson:"success"`
		LoggedIn bool `json:"loggedin" bson:"loggedin"`
	}

)
