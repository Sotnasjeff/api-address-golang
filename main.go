package main

import (
	"api-address-golang/repository/configuration"
	"api-address-golang/router"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configuration.Load()
	if err != nil {
		panic(err)
	}

	controller := chi.NewRouter()
	controller.Post("/", router.CreateAddress)
	controller.Get("/", router.GetAllAddress)
	controller.Get("/{id}", router.GetAddressById)
	controller.Put("/{id}", router.UpdateAddress)
	controller.Delete("/{id}", router.DeleteAddress)

	http.ListenAndServe(fmt.Sprintf(":%s", configuration.GetServerPort()), controller)

}
