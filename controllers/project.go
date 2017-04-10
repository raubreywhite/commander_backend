package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"statsconsultant/models"
)


// CreateUser creates a new user resource
func (uc UserController) CreateProject(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// Stub an user to be populated from the body
u := models.Project{}

// Populate the user data
json.NewDecoder(r.Body).Decode(&u)

// Add an Id
u.Id = bson.NewObjectId()

// Write the user to mongo
uc.session.DB("go_rest_tutorial").C("projects").Insert(u)

// Marshal provided interface into JSON structure
uj, _ := json.Marshal(u)

// Write content-type, statuscode, payload
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(201)
fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) EditProject(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// Stub an user to be populated from the body
u := models.Project{}

// Populate the user data
json.NewDecoder(r.Body).Decode(&u)

change := mgo.Change{
Update: u,
ReturnNew: true,
}
fmt.Println("Before")
fmt.Println(u)
_, err := uc.session.DB("go_rest_tutorial").C("projects").FindId(u.Id).Apply(change, &u)
if(err!=nil){
 fmt.Println(err)
} else {
fmt.Println("After")
fmt.Println(u)
}

// Marshal provided interface into JSON structure
uj, _ := json.Marshal(u)

// Write content-type, statuscode, payload
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(201)
fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) GetProject(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// Grab id
u := models.Project{}

// Populate the user data
json.NewDecoder(r.Body).Decode(&u)

id := u.Id

// Fetch user
if err := uc.session.DB("go_rest_tutorial").C("projects").FindId(id).One(&u); err != nil {
fmt.Println(err)
w.WriteHeader(404)
return
}
// Marshal provided interface into JSON structure
uj, _ := json.Marshal(u)

// Write content-type, statuscode, payload
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(200)
fmt.Fprintf(w, "%s", uj)
}
