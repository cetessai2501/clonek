package main

import (
	"fmt"

	"github.com/cetessai2501/clonek/pkg/adding"
	"github.com/cetessai2501/clonek/pkg/reviewing"
	"github.com/cetessai2501/clonek/pkg/storage/json"
)

func main() {

	var adder adding.Service
	var reviewer reviewing.Service

	// error handling omitted for simplicity
	s, _ := json.NewStorage()

	adder = adding.NewService(s)
	reviewer = reviewing.NewService(s)

	// add some sample data
	adder.AddSampleBeers(DefaultBeers)
	reviewer.AddSampleReviews(DefaultReviews)

	fmt.Println("Finished adding sample data.")
}
