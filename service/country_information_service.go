package service

import (
	"country_information_api/domain"
	"country_information_api/dto"
	"country_information_api/repository"
)

type ICountryService interface {
	GetAllCountries() []domain.Country
	GetCountryById(countryId int64) (domain.Country, error)
	GetCountryByName(name string) (domain.Country, error)
	GetCountryByIso2(iso2 string) (domain.Country, error)
	GetCountryByIso3(iso3 string) (domain.Country, error)
	GetCountryByPhoneCode(phoneCode int) (domain.Country, error)
	AddCountry(country dto.CountryDto) error
	UpdateCountryById(country dto.CountryDto, countryId int64) error
	DeleteCountryById(countryId int64) error
}

type CountryService struct {
	countryInformationRepository repository.ICountryInformationRepository
}

func NewCountryService(cir repository.ICountryInformationRepository) ICountryService {
	return &CountryService{
		countryInformationRepository: cir,
	}
}

func (c CountryService) GetAllCountries() []domain.Country {
	return c.countryInformationRepository.GetAllCountries()
}

func (c CountryService) GetCountryById(countryId int64) (domain.Country, error) {
	return c.GetCountryById(countryId)
}

func (c CountryService) GetCountryByName(name string) (domain.Country, error) {
	return c.countryInformationRepository.GetCountryByName(name)
}

func (c CountryService) GetCountryByIso2(iso2 string) (domain.Country, error) {
	return c.countryInformationRepository.GetCountryByIso2(iso2)
}

func (c CountryService) GetCountryByIso3(iso3 string) (domain.Country, error) {
	return c.countryInformationRepository.GetCountryByIso3(iso3)
}

func (c CountryService) GetCountryByPhoneCode(phoneCode int) (domain.Country, error) {
	return c.countryInformationRepository.GetCountryByPhoneCode(phoneCode)
}

func (c CountryService) AddCountry(country dto.CountryDto) error {
	return c.countryInformationRepository.AddCountry(domain.Country{
		Iso:       country.Iso,
		Name:      country.Name,
		NiceName:  country.NiceName,
		Iso3:      country.Iso3,
		Numcode:   country.Numcode,
		PhoneCode: country.PhoneCode,
	})
}

func (c CountryService) UpdateCountryById(country dto.CountryDto, countryId int64) error {
	return c.countryInformationRepository.UpdateCountryById(domain.Country{
		Iso:       country.Iso,
		Name:      country.Name,
		NiceName:  country.NiceName,
		Iso3:      country.Iso3,
		Numcode:   country.Numcode,
		PhoneCode: country.PhoneCode,
	}, countryId)
}

func (c CountryService) DeleteCountryById(countryId int64) error {
	return c.countryInformationRepository.DeleteCountryById(countryId)
}
