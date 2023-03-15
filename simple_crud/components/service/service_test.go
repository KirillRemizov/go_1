package service

import (
	"reflect"
	"testing"

	"github.com/KirillRemizov/GO_1/simple_crud/components/storage"
)

func TestCreate(t *testing.T) {

	str := storage.NewStorage()
	s := NewCrudService(str)

	temperature := 22.1
	windSpeed := 1.2

	result, err := s.CreateWeatherCondition(temperature, windSpeed)
	if err != nil {
		panic(err)
	}

	if result == nil {
		t.Error("result is nil")
		return
	}

	if result.Temperature != temperature {
		t.Error("wrong temperature")
		return
	}

	if result.WindSpeed != windSpeed {
		t.Error("wrong wind speed")
		return
	}

	if result.ID == "" {
		t.Error("id is empty")
		return

	}

	_, err = s.storage.ReadCondition(result.ID)

	if err != nil {
		t.Error("not found in local storage")
	}

}

func TestUpdate(t *testing.T) {

	temperature := 22.1
	windSpeed := 1.2

	str := storage.NewStorage()
	s := NewCrudService(str)

	result, err := s.CreateWeatherCondition(temperature, windSpeed)
	if err != nil {
		panic(err)
	}

	newTemp := temperature + 1
	newWind := windSpeed + 1

	s.UpdateWeatherCondition(result.ID, nil, nil)

	if result.Temperature != temperature || result.WindSpeed != windSpeed {
		t.Error("temperature or winspeed is updated with nil arguments")
		return
	}

	s.UpdateWeatherCondition(result.ID, &newTemp, &newWind)

	if result.Temperature != newTemp && result.WindSpeed != newWind {
		t.Error("temperature and windspeed not updated")
		return

	}

}

func TestDelete(t *testing.T) {

	str := storage.NewStorage()
	s := NewCrudService(str)
	temperature := 22.1
	windSpeed := 1.2

	result, err := s.CreateWeatherCondition(temperature, windSpeed)
	if err != nil {
		panic(err)
	}

	s.DeleteWeatherCondition(result.ID)

	readed, err := s.ReadWeatherCondition(result.ID)

	if readed != nil {
		t.Error("Readed object is no nill")
	}

	if err == nil {
		t.Error("exist in local storage after deletion")
		return
	}

}

func TestRead(t *testing.T) {

	str := storage.NewStorage()
	s := NewCrudService(str)

	temperature := 22.1
	windSpeed := 1.2

	created, err := s.CreateWeatherCondition(temperature, windSpeed)
	if err != nil {
		panic(err)
	}

	readed, error := s.ReadWeatherCondition(created.ID)
	if error != nil {
		t.Error("read method returns error")
		return
	}

	equal := reflect.DeepEqual(created, readed)
	if !equal {
		t.Error("not equal")
		return
	}

}

func TestList(t *testing.T) {

	str := storage.NewStorage()
	s := NewCrudService(str)

	temperature := 22.1
	windSpeed := 1.2

	ids := make(map[string]struct{}) // map to store added id

	//adding three new conditions
	first, err := s.CreateWeatherCondition(temperature, windSpeed)
	if err != nil {
		panic(err)
	}

	second, err := s.CreateWeatherCondition(temperature, windSpeed)
	if err != nil {
		panic(err)
	}

	third, err := s.CreateWeatherCondition(temperature, windSpeed)
	if err != nil {
		panic(err)
	}

	// adding new conditions's id to map
	ids[first.ID] = struct{}{}
	ids[second.ID] = struct{}{}
	ids[third.ID] = struct{}{}

	// getting list
	list, err := s.ListWeatherConditions()

	if err != nil {
		t.Error("list method returns error")
		return
	}

	for _, v := range list {
		_, exist := ids[v.ID]
		if !exist {
			t.Error("id is not in list")
			return
		}
	}

	// deleting first condition
	s.DeleteWeatherCondition(first.ID)

	// getting new list
	list, _ = s.ListWeatherConditions()

	// Checking if deleted condition is in the list
	for _, v := range list {
		if v.ID == first.ID {
			t.Error("Deleted id of first is in list")
			return
		}
	}

	// deleting third condition
	s.DeleteWeatherCondition(third.ID)

	// getting new list
	list, _ = s.ListWeatherConditions()

	for _, v := range list {

		if v.ID == third.ID {
			t.Error("Deleted id of third in the list ")
		}
	}

	// checkin if second is still in the list
	for _, v := range list {
		if v.ID == second.ID {
			return
		}
	}

	t.Error("Second is not in the list")

}
