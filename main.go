package main

// curl -H "Content-Type: application/json" -X POST -d '{"id":"58b315e935d54400019a6c79", "name":"xxx", "statistician":"58b314c253ae1e0001bdab86","client":"58b314c253ae1e0001bdab86"}' http://localhost:8080/getproject

// curl -H "Content-Type: application/json" -X POST -d '{"type":"hello", "name":"xxx"}' http://localhost:8080/create/users

import (
	// Standard library packages
	"fmt"
	"net/http"
  "log"

	// Third party packages
  "github.com/kardianos/service"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"github.com/raubreywhite/commander_backend/controllers"
	"github.com/rs/cors"

)


var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
  if service.Interactive() {
		logger.Info("Running in terminal.")
	} else {
		logger.Info("Running under service manager.")
	}

	go p.run()
	return nil
}
func (p *program) run() {
	// Instantiate a new router
	r := httprouter.New()

	// Get a UserController instance
	session := getSession()
	uc := controllers.NewUserController(session)
	defer session.Close()

	// login
	r.POST("/login", uc.Login)

	r.POST("/edit/:type", uc.Edit)

	r.POST("/create/:type", uc.Create)

	r.POST("/delete/:type", uc.Delete)

	r.POST("/get/:type", uc.Get)

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
	})

	// Fire up the server
	http.ListenAndServe(":8080", c.Handler(r))
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
  logger.Info("I'm Stopping!")
	return nil
}

func main() {
  svcConfig := &service.Config{
		Name:        "GoServiceTest",
		DisplayName: "Go Service Test",
		Description: "This is a test Go service.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	ip := "mongodb://mongo"

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
