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

	srv.UpdateWeatherCondition(condition.ID, 25.3, 0.11)

	srv.ListLocalStorage()

	fmt.Println("Read Condition:")
	spew.Dump(srv.ReadWeatherCondition(condition.ID))

	srv.DeleteWeatherCondition(condition.ID)

	srv.ListLocalStorage()
}
