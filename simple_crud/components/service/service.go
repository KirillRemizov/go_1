package service

import (
	"errors"

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

	/* Creates condition object.
	   Adds pointer of new object to map storage. */

	condition := &types.WeatherCondition{
		ID:          uuid.NewV4().String(),
		Temperature: temperature,
		WindSpeed:   windSpeed,
	}

	s.localStorage[condition.ID] = condition

	return condition, nil
}

func (s *CrudService) DeleteWeatherCondition(id string) error {

	/* Deletes pointer of condition from map storage by id.
	   Returns error if not found.
	   What will happen with the object? */

	_, exists := s.localStorage[id]
	if exists {
		delete(s.localStorage, id)
		return nil
	}
	return errors.New("not found")

}

func (s *CrudService) UpdateWeatherCondition(id string, temperature *float64, windSpeed *float64) error {

	condition, exists := s.localStorage[id]

	if !exists {
		return errors.New("not found")

	}
	if temperature != nil {
		condition.Temperature = *temperature
	}
	if windSpeed != nil {
		condition.WindSpeed = *windSpeed
	}

	return nil

}

func (s *CrudService) ReadWeatherCondition(id string) (*types.WeatherCondition, error) {
	// Returns pointer to condition object if it exist in map

	condition, exists := s.localStorage[id]

	if !exists {
		return nil, errors.New("not found")
	}

	return condition, nil
}

func (s *CrudService) ListWeatherConditions() ([]*types.WeatherCondition, error) {
	//
	list := []*types.WeatherCondition{}

	for _, k := range s.localStorage {
		list = append(list, k)
	}

	return list, nil

}
