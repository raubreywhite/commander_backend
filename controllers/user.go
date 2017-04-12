package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/raubreywhite/commander_backend/models"
)

type (
	// UserController represents the controller for operating on the User resource
	UserController struct {
		session *mgo.Session
	}
)

// NewUserController provides a reference to a UserController with provided mongo session
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// GetUser retrieves an individual user resource
func (uc UserController) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	u := models.User{}
	//uFail := models.User{Success:false, LoggedIn:false}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&u)

	// Grab id
	email := u.Email
	password := u.Password

	// Verify id is ObjectId, otherwise bail
	if email == ""  || password==""{
		w.WriteHeader(404)
		return
	}

	// Stub user
	token, err := GenerateRandomString(32)
	if err != nil {
	// Serve an appropriately vague error to the
	// user, but log the details internally.
	}
	//us := models.UserSession{Id:bson.NewObjectId(), Session:token, Success:false}
	fmt.Println(email)
	fmt.Println(password)
	time.Sleep(1000)
	// Fetch user
	if err := uc.session.DB("go_rest_tutorial").C("users").Find(bson.M{"email": email}).One(&u); err == nil {
		fmt.Println("FOUND")
		fmt.Println(u)
		if password==u.Password {
			fmt.Println("Before")
			fmt.Println(u)

			u.Session = token
			u.Success = true
			u.LoggedIn = true

			change := mgo.Change{
				Update: bson.M{"$set": bson.M{"session": u.Session, "success" : u.Success, "loggedin": u.LoggedIn}},
				ReturnNew: true,
				}

			_, err = uc.session.DB("go_rest_tutorial").C("users").FindId(u.Id).Apply(change, &u)
			fmt.Println("After")
			fmt.Println(u)

		} else {
			fmt.Println("WRONG PW")
			u = models.User{Success:false, LoggedIn:false}
		}
	} else {
		fmt.Println(err)
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}



// CreateUser creates a new user resource
func (uc UserController) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

// Stub an user to be populated from the body
structU := models.User{}
structP := models.Project{}
structE := models.Entry{}

id := bson.NewObjectId()
var f interface{}

// Populate the user data
structType := p.ByName("type")
if(structType=="users"){
json.NewDecoder(r.Body).Decode(&structU)
structU.Id = id
f = structU
} else if(structType=="projects"){
json.NewDecoder(r.Body).Decode(&structP)
structP.Id = id
fmt.Println(structP)
f = structP
} else if(structType=="entries"){
json.NewDecoder(r.Body).Decode(&structE)
structE.Id = id
f = structE
}

fmt.Println(f)

// Write the user to mongo
err := uc.session.DB("go_rest_tutorial").C(structType).Insert(f)
if err != nil {
fmt.Println(err)
}

// Marshal provided interface into JSON structure
uj, _ := json.Marshal(f)

// Write content-type, statuscode, payload
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(201)
fmt.Fprintf(w, "%s", uj)
}


func (uc UserController) Edit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

// Stub an user to be populated from the body
structU := models.User{}
structP := models.Project{}
structE := models.Entry{}

id := bson.NewObjectId()
var f interface{}

// Populate the user data
structType := p.ByName("type")
if(structType=="users"){
json.NewDecoder(r.Body).Decode(&structU)
id = structU.Id
f = structU
} else if(structType=="projects"){
json.NewDecoder(r.Body).Decode(&structP)
id = structP.Id
f = structP
} else if(structType=="entries"){
json.NewDecoder(r.Body).Decode(&structE)
id = structE.Id
f = structE
}
fmt.Println(id)
change := mgo.Change{
Update: f,
ReturnNew: true,
}

_, err := uc.session.DB("go_rest_tutorial").C(structType).FindId(id).Apply(change, &f)
if(err!=nil){
fmt.Println(err)
} else {
fmt.Println("After")
fmt.Println(f)
}

// Marshal provided interface into JSON structure
uj, _ := json.Marshal(f)

// Write content-type, statuscode, payload
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(201)
fmt.Fprintf(w, "%s", uj)
}



// RemoveUser removes an existing user resource
func (uc UserController) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// Grab id
id := p.ByName("id")

// Verify id is ObjectId, otherwise bail
if !bson.IsObjectIdHex(id) {
w.WriteHeader(404)
return
}

// Grab id
oid := bson.ObjectIdHex(id)

// Remove user
if err := uc.session.DB("go_rest_tutorial").C("users").RemoveId(oid); err != nil {
w.WriteHeader(404)
return
}

// Write status
w.WriteHeader(200)
}


func (uc UserController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

// Stub an user to be populated from the body
structU := models.User{}
structP := models.Project{}
structE := models.Entry{}

//id := bson.NewObjectId()
var f interface{}

// Populate the user data
structType := p.ByName("type")
if(structType=="users"){
json.NewDecoder(r.Body).Decode(&structU)
	if err := uc.session.DB("go_rest_tutorial").C(structType).FindId(structU.Id).One(&structU); err != nil {
		w.WriteHeader(404)
		return
	}
f = structU
} else if(structType=="projects"){
json.NewDecoder(r.Body).Decode(&structP)
	if err := uc.session.DB("go_rest_tutorial").C(structType).FindId(structP.Id).One(&structP); err != nil {
		w.WriteHeader(404)
		return
	}
f = structP
} else if(structType=="entries"){
json.NewDecoder(r.Body).Decode(&structE)
	if err := uc.session.DB("go_rest_tutorial").C(structType).FindId(structE.Id).One(&structE); err != nil {
		w.WriteHeader(404)
		return
	}
f = structE
}

// Marshal provided interface into JSON structure
uj, _ := json.Marshal(f)

// Write content-type, statuscode, payload
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(200)
fmt.Fprintf(w, "%s", uj)
}















// GetUser retrieves an individual user resource
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub user
	u := models.User{}
fmt.Println(oid)
fmt.Println(id)
fmt.Println(1)

	// Fetch user
	if err := uc.session.DB("go_rest_tutorial").C("users").FindId(id).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
fmt.Println(2)
	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// CreateUser creates a new user resource
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an user to be populated from the body
	u := models.User{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&u)

	// Add an Id
	u.Id = bson.NewObjectId()

	// Write the user to mongo
	uc.session.DB("go_rest_tutorial").C("users").Insert(u)

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

// RemoveUser removes an existing user resource
func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err := uc.session.DB("go_rest_tutorial").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// Write status
	w.WriteHeader(200)
}
