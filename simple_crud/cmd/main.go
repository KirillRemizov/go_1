package main

import (
	"fmt"
	"log"

	"github.com/KirillRemizov/GO_1/simple_crud/components/service"
	"github.com/davecgh/go-spew/spew"
)

func main() {

	log.SetFlags(log.Lshortfile)

	srv := service.NewCrudService()

	condition, err := srv.CreateWeatherCondition(22.2, 0.73)
	if err != nil {
		log.Println(err)
	}

	spew.Dump(condition)
	temp, wind := 25.3, 0.11
	srv.UpdateWeatherCondition(condition.ID, &temp, &wind)

	srv.ListWeatherConditions()

	fmt.Println("Read Condition:")
	spew.Dump(srv.ReadWeatherCondition(condition.ID))

	fmt.Println("Printing list:")
	fmt.Println(srv.ListWeatherConditions())

	srv.DeleteWeatherCondition(condition.ID)

	fmt.Println("Printing list after deletion:")
	fmt.Println(srv.ListWeatherConditions())

}
