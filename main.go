package main

// curl -H "Content-Type: application/json" -X POST -d '{"id":"58b315e935d54400019a6c79", "name":"xxx", "statistician":"58b314c253ae1e0001bdab86","client":"58b314c253ae1e0001bdab86"}' http://localhost:8080/getproject

// curl -H "Content-Type: application/json" -X POST -d '{"type":"hello", "name":"xxx"}' http://localhost:8080/create/users

import (
	// Standard library packages
	"fmt"
	"net/http"
	"os"

	// Third party packages
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"github.com/raubreywhite/commander_backend/controllers"
	"github.com/raubreywhite/commander_backend/models"

	"gopkg.in/mgo.v2/bson"
	"github.com/rs/cors"

)

func main() {

	// Instantiate a new router
	r := httprouter.New()

	// Get a UserController instance
	session := getSession()
	uc := controllers.NewUserController(session)
	defer session.Close()


u := models.User{Id: bson.NewObjectId(), Email: "r@rwhite.no", Password: "hello"}
err := session.DB("go_rest_tutorial").C("users").Insert(u)
fmt.Println(err)
fmt.Println(u)


	// login
	r.POST("/login", uc.Login)

	r.POST("/edit/:type", uc.Edit)

	r.POST("/create/:type", uc.Create)

	r.POST("/delete/:type", uc.Delete)

	r.POST("/get/:type", uc.Get)

	// Get a user resource
	r.GET("/user/:id", uc.GetUser)

	// Create a new user
	r.POST("/user", uc.CreateUser)

	// Remove an existing user
	r.DELETE("/user/:id", uc.RemoveUser)

	// Create a new project
	r.POST("/createproject", uc.CreateProject)

// Edit a project
r.POST("/editproject", uc.EditProject)

// Get a project
r.POST("/getproject", uc.GetProject)

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
	})

	// Fire up the server
	http.ListenAndServe(":8080", c.Handler(r))
}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	ip := "mongodb://localhost"
	if os.Getenv("HOSTIP") != "" {
		ip = "mongodb://" + os.Getenv("HOSTIP")
	}
	fmt.Println(ip)

	s, err := mgo.Dial(ip)

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

err = s.DB("go_rest_tutorial").DropDatabase()
if err != nil {
panic(err)
}

	// Deliver session
	return s
}
