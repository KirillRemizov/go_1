package main

import (
	"fmt"
	"log"

	"github.com/KirillRemizov/GO_1/simple_crud/components/service"
	"github.com/KirillRemizov/GO_1/simple_crud/components/storage"
	"github.com/davecgh/go-spew/spew"
)

func main() {

	log.SetFlags(log.Lshortfile)

	str := storage.NewStorage()
	srv := service.NewCrudService(str)

	condition, err := srv.CreateWeatherCondition(22.2, 0.73)
	if err != nil {
		log.Println(err)
	}

	condition_2, err := srv.CreateWeatherCondition(30.2, 0.9)

	if err != nil {
		log.Println(err)
	}
	spew.Dump(condition)
	spew.Dump(condition_2)

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
