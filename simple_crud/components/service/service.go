package service

import (
	"errors"

	"github.com/KirillRemizov/GO_1/simple_crud/types"
	uuid "github.com/satori/go.uuid"
)

type CrudService struct {
	storage types.Storage
}

func NewCrudService(storage types.Storage) *CrudService {
	s := &CrudService{storage: storage}
	return s
}

//var _ types.Service = (*CrudService)(nil)

func (s *CrudService) CreateWeatherCondition(temperature float64, windSpeed float64) (*types.WeatherCondition, error) {

	// Creates condition object.

	condition := &types.WeatherCondition{
		ID:          uuid.NewV4().String(),
		Temperature: temperature,
		WindSpeed:   windSpeed,
	}

	err := s.storage.StoreCondition(condition)

	if err != nil {
		return condition, errors.New("unable to store condition")
	}

	return condition, nil
}

func (s *CrudService) DeleteWeatherCondition(id string) error {

	return s.storage.DeleteCondition(id)

}

func (s *CrudService) UpdateWeatherCondition(id string, temperature *float64, windSpeed *float64) error {

	err := s.storage.UpdateCondition(id, temperature, windSpeed)

	if err != nil {
		return err
	}
	return nil
}

func (s *CrudService) ReadWeatherCondition(id string) (*types.WeatherCondition, error) {
	// Returns pointer to condition object if it exist in map

	condition, err := s.storage.ReadCondition(id)

	if err != nil {
		return nil, err
	}

	return condition, nil
}

func (s *CrudService) ListWeatherConditions() ([]*types.WeatherCondition, error) {

	list, err := s.storage.ListConditions()

	if err != nil {
		return nil, err
	}

	return list, nil
}
