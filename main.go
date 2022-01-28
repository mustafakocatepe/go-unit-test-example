package main

import (
	"fmt"

	"github.com/mustafakocatepe/go-unit-test-example/control"
	"github.com/mustafakocatepe/go-unit-test-example/model"
)

func main() {
	var shelter = model.NewShelter()

	response, err := control.ControlMedicalStatus(&shelter)
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
