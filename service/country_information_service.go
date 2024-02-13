package service

import (
	"country_information_api/dto/response"
	"country_information_api/repository"
)

type ICountryService interface {
	GetAllCountries() []response.CountryResponseDto
	GetCountryById(countryId int64) (response.CountryResponseDto, error)
	GetCountryByName(name string) (response.CountryResponseDto, error)
	GetCountryByIso2(iso2 string) (response.CountryResponseDto, error)
	GetCountryByIso3(iso3 string) (response.CountryResponseDto, error)
	GetCountryByPhoneCode(phoneCode int) (response.CountryResponseDto, error)
}

type CountryService struct {
	countryInformationRepository repository.ICountryInformationRepository
}

func NewCountryService(cir repository.ICountryInformationRepository) ICountryService {
	return &CountryService{
		countryInformationRepository: cir,
	}
}

func (c *CountryService) GetAllCountries() []response.CountryResponseDto {
	return c.countryInformationRepository.GetAllCountries()
}

func (c *CountryService) GetCountryById(countryId int64) (response.CountryResponseDto, error) {
	return c.countryInformationRepository.GetCountryById(countryId)
}

func (c *CountryService) GetCountryByName(name string) (response.CountryResponseDto, error) {
	return c.countryInformationRepository.GetCountryByName(name)
}

func (c *CountryService) GetCountryByIso2(iso2 string) (response.CountryResponseDto, error) {
	return c.countryInformationRepository.GetCountryByIso2(iso2)
}

func (c *CountryService) GetCountryByIso3(iso3 string) (response.CountryResponseDto, error) {
	return c.countryInformationRepository.GetCountryByIso3(iso3)
}

func (c *CountryService) GetCountryByPhoneCode(phoneCode int) (response.CountryResponseDto, error) {
	return c.countryInformationRepository.GetCountryByPhoneCode(phoneCode)
}
