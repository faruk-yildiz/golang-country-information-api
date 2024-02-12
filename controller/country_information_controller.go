package controller

import (
	"country_information_api/dto"
	"country_information_api/dto/response"
	"country_information_api/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CountryController struct {
	countryService service.ICountryService
}

func NewCountryController(countryService service.ICountryService) *CountryController {
	return &CountryController{
		countryService: countryService,
	}
}

func (countryController *CountryController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/countries", countryController.GetAllCountries)
	e.GET("/api/v1/countries/id", countryController.GetCountryById)
	e.GET("/api/v1/countries/name", countryController.GetCountryByName)
	e.GET("/api/v1/countries/iso2", countryController.GetCountryByIso2)
	e.GET("/api/v1/countries/iso3", countryController.GetCountryByIso3)
	e.GET("/api/v1/countries/phonecode", countryController.GetCountryByPhoneCode)
	e.POST("/api/v1/countries", countryController.AddCountry)
	e.PUT("/api/v1/countries/:id", countryController.UpdateCountry)
	e.DELETE("/api/v1/countries/:id", countryController.DeleteProduct)
}

func (countryController *CountryController) GetAllCountries(e echo.Context) error {
	countries := countryController.countryService.GetAllCountries()
	return e.JSON(http.StatusOK, countries)
}

func (countryController *CountryController) GetCountryById(e echo.Context) error {
	param := e.QueryParam("id")
	countryId, err := strconv.Atoi(param)

	country, err := countryController.countryService.GetCountryById(int64(countryId))

	if err != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	return e.JSON(http.StatusOK, country)
}

func (countryController *CountryController) GetCountryByName(e echo.Context) error {
	param := e.QueryParam("name")

	country, err := countryController.countryService.GetCountryByName(param)

	if err != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	return e.JSON(http.StatusOK, country)
}

func (countryController *CountryController) GetCountryByIso2(e echo.Context) error {
	param := e.QueryParam("iso2")

	country, err := countryController.countryService.GetCountryByIso2(param)

	if err != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	return e.JSON(http.StatusOK, country)
}
func (countryController *CountryController) GetCountryByIso3(e echo.Context) error {
	param := e.QueryParam("iso3")

	country, err := countryController.countryService.GetCountryByIso3(param)

	if err != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	return e.JSON(http.StatusOK, country)
}

func (countryController *CountryController) GetCountryByPhoneCode(e echo.Context) error {
	param := e.QueryParam("phonecode")
	phoneCode, convertError := strconv.Atoi(param)

	if convertError != nil {
		return e.NoContent(http.StatusBadRequest)
	}

	country, err := countryController.countryService.GetCountryByPhoneCode(phoneCode)

	if err != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: err.Error()})
	}

	return e.JSON(http.StatusOK, country)
}

func (countryController *CountryController) AddCountry(e echo.Context) error {
	var addCountryRequestObject dto.UpdateOrAddCountryDto
	err := e.Bind(&addCountryRequestObject)

	if err != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: err.Error()})
	}
	addCountryErr := countryController.countryService.AddCountry(addCountryRequestObject)

	if addCountryErr != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: err.Error()})
	}
	return e.NoContent(http.StatusCreated)
}

func (countryController *CountryController) UpdateCountry(e echo.Context) error {
	param := e.Param("id")
	countryIdInt, convertError := strconv.Atoi(param)
	countryId := int64(countryIdInt)

	if convertError != nil {
		return e.NoContent(http.StatusBadRequest)
	}

	var addCountryRequestObject dto.UpdateOrAddCountryDto
	err := e.Bind(&addCountryRequestObject)

	if err != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: err.Error()})
	}
	updateCountryError := countryController.countryService.UpdateCountryById(addCountryRequestObject, countryId)

	if updateCountryError != nil {
		return e.JSON(http.StatusBadRequest, response.ErrorResponse{ErrorDescription: err.Error()})
	}
	return e.NoContent(http.StatusOK)
}

func (countryController *CountryController) DeleteProduct(e echo.Context) error {
	param := e.Param("id")
	atoi, _ := strconv.Atoi(param)
	countryId := int64(atoi)

	err := countryController.countryService.DeleteCountryById(countryId)

	if err != nil {
		return e.JSON(http.StatusNotFound, response.ErrorResponse{err.Error()})
	}
	return e.NoContent(http.StatusOK)
}
