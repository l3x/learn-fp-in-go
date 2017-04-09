package main

import (
	. "github.com/l3x/learn-fp-in-go/chapter4/01_hof"
	"log"
	"os"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func init() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)
}

func main() {
	if os.Getenv("RUN_HTTP_SERVER") == "TRUE" {
		router := httprouter.New()
		router.GET("/cars", CarsIndexHandler)
		router.GET("/cars/:id", CarHandler)
		log.Println("Listening on port 8000")
		log.Fatal(http.ListenAndServe(":8000", router))
	} else {
		cars := LoadCars()

		PrintCars("ByMake - Honda", cars.Filter(ByMake("Honda")))

		PrintCars("Numeric", cars.Filter(ByHasNumber()))

		PrintCars("Foreign, Numeric, Toyota",
			cars.Filter(ByForeign()).
				Filter(ByHasNumber()).
				Filter(ByMake("Toyota")))

		moreCars := LoadMoreCars()

		PrintCars("More Cars, Domestic, Numeric, GM",
			cars.AddCars(moreCars).
				Filter(ByDomestic()).
				Filter(ByHasNumber()).
				Filter(ByMake("GM")))

		PrintCars("Numeric, Foreign, Map Upgraded",
			cars.Filter(ByHasNumber()).
				Filter(ByForeign()).
				Map(Upgrade()))

		PrintCars("Filter Honda, Reduce JSON",
			cars.Filter(ByMake("Honda")).
				Reduce(JsonReducer(cars), Collection{}))

		PrintCars("Reduce, Honda, JSON",
			cars.Reduce(MakeReducer("Honda", cars), Collection{}).
				Reduce(JsonReducer(cars), Collection{}))

		PrintCars2("Reduce - Lexus",
			cars.Filter(ByMake("Lexus")).
				Reduce2(CarTypeReducer(cars), []CarType{}))

		PrintCars("ByModel - Accord up/downgraded",
			cars.Filter(ByModel("Accord")).
				Map(Upgrade()).
				Map(Downgrade()))

		PrintCars("GenerateCars(1, 3)",
			cars.GenerateCars(1, 3))

		PrintCars("GenerateCars(1, 14), Domestic, Numeric, JSON",
			cars.GenerateCars(1, 14).
				Filter(ByDomestic()).
				Map(Upgrade()).
				Filter(ByHasNumber()).
				Reduce(JsonReducer(cars), Collection{}))

	}
}

