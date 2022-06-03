package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"go-test-areas/app"
	"go-test-areas/db"
	ihttp "go-test-areas/http"
	"go-test-areas/mocks"
	"log"
	"net/http"
	"os"
)

func main() {
	args := os.Args
	op := "server"
	if len(args) > 1 {
		op = args[0]
	}
	if err := run(op); err != nil {
		fmt.Println(fmt.Errorf("error - server failed to start. err: %v", err))
	}
}

func run(op string) error {
	if op == "http_test" {
		// mocking the pet svc to test our http routes
		svc := mocks.PetSvc{}
		h := ihttp.NewHandler(svc)
		r := chi.NewRouter()
		ihttp.Routes(r, h)
		return http.ListenAndServe(":8080", r)
	}

	// tying up all the components together and running the server
	d, err := db.NewMongoStore()
	if err != nil {
		return errors.Wrap(err, "unable to intialize db")
	}
	svc := app.NewPetSvc(d)
	h := ihttp.NewHandler(svc)
	r := chi.NewRouter()
	ihttp.Routes(r, h)
	log.Print("all good")
	return http.ListenAndServe(":8080", r)
}
