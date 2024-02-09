package repository

import (
	"context"
	"country_information_api/domain"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"strings"
)

var ctx context.Context

type ICountryInformationRepository interface {
	GetAllCountries() []domain.Country
	GetCountryById(countryId int64) (domain.Country, error)
	GetCountryByName(name string) (domain.Country, error)
	GetCountryByIso2(iso2 string) (domain.Country, error)
	GetCountryByIso3(iso3 string) (domain.Country, error)
	GetCountryByPhoneCode(phoneCode int) (domain.Country, error)
	AddCountry(country domain.Country) error
	UpdateCountryById(country domain.Country, countryId int64) error
	DeleteCountryById(countryId int64) error
}

type CountryInformationRepository struct {
	dbPool *pgxpool.Pool
}

func NewCountryInformationRepository(dbpool *pgxpool.Pool) ICountryInformationRepository {
	return &CountryInformationRepository{
		dbPool: dbpool,
	}
}

func (c CountryInformationRepository) GetAllCountries() []domain.Country {
	ctx = context.Background()

	rows, err := c.dbPool.Query(ctx, "Select * from countries")
	if err != nil {
		log.Errorf("Error while getting all countries %v", err)
		return []domain.Country{}
	}
	var countries []domain.Country
	var country domain.Country
	for rows.Next() {
		err := rows.Scan(&country.Id, &country.Iso, &country.Name, &country.NiceName, &country.Iso3, &country.Numcode, &country.PhoneCode)
		if err != nil {
			return nil
		}
		countries = append(countries, country)
	}
	return countries

}

func (c CountryInformationRepository) GetCountryById(countryId int64) (domain.Country, error) {
	ctx = context.Background()
	getByIdSql := `Select * from countries where id=$1`
	result := c.dbPool.QueryRow(ctx, getByIdSql, countryId)

	var country domain.Country

	err := result.Scan(&country.Id, &country.Iso, &country.Name, &country.NiceName, &country.Iso3, &country.Numcode, &country.PhoneCode)
	if err != nil {
		return domain.Country{}, errors.New(fmt.Sprintf("Error while getting product with id %v", countryId))
	}

	return country, nil
}

func (c CountryInformationRepository) GetCountryByName(name string) (domain.Country, error) {
	ctx = context.Background()
	name = strings.ToUpper(name)
	getByNameSql := `Select * from countries where name=$1`
	var country domain.Country

	result := c.dbPool.QueryRow(ctx, getByNameSql, name)
	err := result.Scan(&country.Id, &country.Iso, &country.Name, &country.NiceName, &country.Iso3, &country.Numcode, &country.PhoneCode)

	if err != nil {
		return domain.Country{}, errors.New(fmt.Sprintf("Not found country with name %v", name))
	}
	return country, nil
}

func (c CountryInformationRepository) GetCountryByIso2(iso2 string) (domain.Country, error) {
	ctx = context.Background()
	iso2 = strings.ToUpper(iso2)
	getByNameSql := `Select * from countries where iso2=$1`
	var country domain.Country

	result := c.dbPool.QueryRow(ctx, getByNameSql, iso2)
	err := result.Scan(&country.Id, &country.Iso, &country.Name, &country.NiceName, &country.Iso3, &country.Numcode, &country.PhoneCode)

	if err != nil {
		return domain.Country{}, errors.New(fmt.Sprintf("Not found country with iso2 code %v", iso2))
	}
	return country, nil
}

func (c CountryInformationRepository) GetCountryByIso3(iso3 string) (domain.Country, error) {
	ctx = context.Background()
	iso3 = strings.ToUpper(iso3)
	getByNameSql := `Select * from countries where iso3=$1`
	var country domain.Country

	result := c.dbPool.QueryRow(ctx, getByNameSql, iso3)
	err := result.Scan(&country.Id, &country.Iso, &country.Name, &country.NiceName, &country.Iso3, &country.Numcode, &country.PhoneCode)

	if err != nil {
		return domain.Country{}, errors.New(fmt.Sprintf("Not found country with iso3 code %v", iso3))
	}
	return country, nil
}

func (c CountryInformationRepository) GetCountryByPhoneCode(phoneCode int) (domain.Country, error) {
	ctx = context.Background()
	getByNameSql := `Select * from countries where phonecode=$1`
	var country domain.Country

	result := c.dbPool.QueryRow(ctx, getByNameSql, phoneCode)
	err := result.Scan(&country.Id, &country.Iso, &country.Name, &country.NiceName, &country.Iso3, &country.Numcode, &country.PhoneCode)

	if err != nil {
		return domain.Country{}, errors.New(fmt.Sprintf("Not found country with phonecode  %v", phoneCode))
	}
	return country, nil
}

func (c CountryInformationRepository) AddCountry(country domain.Country) error {
	ctx = context.Background()
	addCountrySql := `Insert into countries (iso,name,nicename,iso3,numcode,phonecode) values ($1,$2,$3,$4,&5,&6)`

	addNewCountry, err := c.dbPool.Exec(ctx, addCountrySql, country.Iso, country.Name, country.NiceName, country.Iso3, country.Numcode, country.PhoneCode)
	if err != nil {
		log.Errorf("Error when adding country to table", err)
		return err
	}
	log.Info(fmt.Printf("Country added %v", addNewCountry))
	return nil
}

func (c CountryInformationRepository) UpdateCountryById(country domain.Country, countryId int64) error {
	ctx = context.Background()

	updateSql := `Update countries set iso = $1,name=$2,nicename=&3,iso3=&4,numcode=&5,phonecode=&6, where id=$7`

	_, err := c.dbPool.Exec(ctx, updateSql, country.Iso, strings.ToUpper(country.Name), country.Iso3, country.Numcode, country.PhoneCode, countryId)

	if err != nil {
		return errors.New(fmt.Sprintf("Error while updating with id : %d", countryId))
	}
	log.Info("Product price updated with id %v", countryId)
	return nil
}

func (c CountryInformationRepository) DeleteCountryById(countryId int64) error {
	ctx = context.Background()
	_, err := c.GetCountryById(countryId)

	if err != nil {
		return errors.New("Country Not Found")
	}

	deleteCountrySql := `Delete from countries where id=$1`

	_, err = c.dbPool.Exec(ctx, deleteCountrySql, countryId)

	if err != nil {
		log.Errorf("Error when deleting product ", err)
		return errors.New(fmt.Sprintf("Error while deleting country"))
	}
	log.Info(fmt.Printf("country deleted "))
	return nil
}
