package integration

import (
	"context"
	"country_information_api/common/postgre"
	"country_information_api/dto/response"
	"country_information_api/repository"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var dbPool *pgxpool.Pool
var countryRepository repository.ICountryInformationRepository
var ctx context.Context

func TestMain(t *testing.M) {
	ctx = context.Background()

	dbPool = postgre.GetConnectionPool(ctx, postgre.Config{
		Host:                  "localhost",
		Port:                  "7432",
		DbName:                "countryinfoapp",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "30s",
	})

	countryRepository = repository.NewCountryInformationRepository(dbPool)
	exitCode := t.Run()
	os.Exit(exitCode)

}
func Test_GetAllCountries(t *testing.T) {
	t.Run("GetAllCountries", func(t *testing.T) {
		actualCountries := countryRepository.GetAllCountries()
		assert.Equal(t, 239, len(actualCountries))
	})
}
func Test_GetCountryById(t *testing.T) {
	var id int64 = 2

	expectedCountry := response.CountryResponseDto{
		Id:        2,
		Iso:       "AL",
		Name:      "ALBANIA",
		NiceName:  "Albania",
		Iso3:      "ALB",
		NumCode:   8,
		PhoneCode: 355,
	}
	t.Run("GetCountryById", func(t *testing.T) {
		actualCountry, _ := countryRepository.GetCountryById(id)
		assert.Equal(t, expectedCountry, actualCountry)
	})
}
func Test_GetCountryByName(t *testing.T) {

	var name = "ALBANIA"

	expectedCountry := response.CountryResponseDto{
		Id:        2,
		Iso:       "AL",
		Name:      "ALBANIA",
		NiceName:  "Albania",
		Iso3:      "ALB",
		NumCode:   8,
		PhoneCode: 355,
	}
	t.Run("GetCountryById", func(t *testing.T) {
		actualCountry, _ := countryRepository.GetCountryByName(name)
		assert.Equal(t, expectedCountry, actualCountry)
	})
}
func Test_GetCountryByIso2(t *testing.T) {
	var iso2 = "AL"

	expectedCountry := response.CountryResponseDto{
		Id:        2,
		Iso:       "AL",
		Name:      "ALBANIA",
		NiceName:  "Albania",
		Iso3:      "ALB",
		NumCode:   8,
		PhoneCode: 355,
	}
	t.Run("GetCountryById", func(t *testing.T) {
		actualCountry, _ := countryRepository.GetCountryByIso2(iso2)
		assert.Equal(t, expectedCountry, actualCountry)
	})
}
func Test_GetCountryByIso3(t *testing.T) {
	var iso3 = "ALB"

	expectedCountry := response.CountryResponseDto{
		Id:        2,
		Iso:       "AL",
		Name:      "ALBANIA",
		NiceName:  "Albania",
		Iso3:      "ALB",
		NumCode:   8,
		PhoneCode: 355,
	}
	t.Run("GetCountryById", func(t *testing.T) {
		actualCountry, _ := countryRepository.GetCountryByIso3(iso3)
		assert.Equal(t, expectedCountry, actualCountry)
	})
}
func Test_GetCountryByPhoneCode(t *testing.T) {
	var phonecode int = 355

	expectedCountry := response.CountryResponseDto{
		Id:        2,
		Iso:       "AL",
		Name:      "ALBANIA",
		NiceName:  "Albania",
		Iso3:      "ALB",
		NumCode:   8,
		PhoneCode: 355,
	}
	t.Run("GetCountryById", func(t *testing.T) {
		actualCountry, _ := countryRepository.GetCountryByPhoneCode(phonecode)
		assert.Equal(t, expectedCountry, actualCountry)
	})
}
