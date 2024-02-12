package main

import (
	"context"
	"country_information_api/common/app"
	"country_information_api/common/postgre"
	"country_information_api/controller"
	"country_information_api/repository"
	"country_information_api/service"
	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	e := echo.New()

	manager := app.NewConfigurationManager()

	dbPool := postgre.GetConnectionPool(ctx, manager.PostgreConfig)

	countryRepository := repository.NewCountryInformationRepository(dbPool)

	countryService := service.NewCountryService(countryRepository)

	countryController := controller.NewCountryController(countryService)

	countryController.RegisterRoutes(e)
	e.Start("localhost:8080")
}
