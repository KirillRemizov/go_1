package service

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestCreate(t *testing.T) {

	fmt.Println("AAAAAAAAAAAAA")

	s := NewCrudService()

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
	}

	if result.WindSpeed != windSpeed {
		t.Error("wrong wind speed")
	}

	if result.ID == "" {
		t.Error("id is empty")

	}

	spew.Dump(result)

	t.Error("MOCK")

}
