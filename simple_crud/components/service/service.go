package service

import (
	"github.com/KirillRemizov/GO_1/simple_crud/types"
	uuid "github.com/satori/go.uuid"
)

type CrudService struct{}

func NewCrudService() *CrudService {
	s := &CrudService{}
	return s
}

func (s *CrudService) CreateWeatherCondition(temperature float64, windSpeed float64) (*types.WeatherCondition, error) {

	condition := &types.WeatherCondition{
		ID:          uuid.NewV4().String(),
		Temperature: temperature,
		WindSpeed:   windSpeed,
	}

	return condition, nil
}
