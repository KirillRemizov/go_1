package storage

import (
	"errors"

	"github.com/KirillRemizov/GO_1/simple_crud/types"
)

type Storage struct {
	localStorage map[string]*types.WeatherCondition
}

func (s *Storage) StoreCondition(condition *types.WeatherCondition) error {

	/* Adds pointer of new object to map storage. */
	_, exist := s.localStorage[condition.ID]

	if exist {
		return errors.New("id found in storage")
	}

	s.localStorage[condition.ID] = condition
	return nil
}

// Deletes pointer of condition from map storage by id.
//
//	Returns error if not found.
func (s *Storage) DeleteCondition(id string) error {

	_, exists := s.localStorage[id]
	if exists {
		delete(s.localStorage, id)
		return nil
	}
	return errors.New("not found")

}

func (s *Storage) UpdateCondition(id string, temperature *float64, windSpeed *float64) error {

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

func (s *Storage) ReadCondition(id string) (*types.WeatherCondition, error) {
	// Returns pointer to condition object if it exist in map

	condition, exists := s.localStorage[id]

	if !exists {
		return nil, errors.New("not found")
	}

	return condition, nil
}

func (s *Storage) ListConditions() ([]*types.WeatherCondition, error) {

	list := make([]*types.WeatherCondition, 0, len(s.localStorage))

	for _, k := range s.localStorage {
		list = append(list, k)
	}

	return list, nil

}
