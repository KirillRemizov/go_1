package service

import (
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
