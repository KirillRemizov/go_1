package service

import (
	"fmt"

	"github.com/KirillRemizov/GO_1/simple_crud/types"
	uuid "github.com/satori/go.uuid"
)

type CrudService struct {
	localStorage map[string]*types.WeatherCondition
}

func NewCrudService() *CrudService {
	s := &CrudService{
		localStorage: make(map[string]*types.WeatherCondition),
	}
	return s
}

func (s *CrudService) CreateWeatherCondition(temperature float64, windSpeed float64) (*types.WeatherCondition, error) {

	condition := &types.WeatherCondition{
		ID:          uuid.NewV4().String(),
		Temperature: temperature,
		WindSpeed:   windSpeed,
	}

	s.localStorage[condition.ID] = condition

	return condition, nil
}

func (s *CrudService) DeleteWeatherCondition(id string) error {

	delete(s.localStorage, id)

	return nil

}

func (s *CrudService) UpdateWeatherCondition(id string, temperature float64, windSpeed float64) error {

	s.localStorage[id].Temperature = temperature
	s.localStorage[id].WindSpeed = windSpeed

	return nil
}

func (s *CrudService) ReadWeatherCondition(id string) (float64, float64) {

	temp, wind := s.localStorage[id].Temperature, s.localStorage[id].WindSpeed

	return temp, wind
}

func (s *CrudService) ListLocalStorage() error {
	fmt.Println("Storage:", s.localStorage)
	return nil
}
