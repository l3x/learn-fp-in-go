package hof

import (
	"sync"
	"log"
)

func carGenerator(iterator func(int) int, lower int, upper int) func() (int, bool) {
	return func() (int, bool) {
		lower = iterator(lower)
		return lower, lower > upper
	}
}

func iterator(i int) int {
	i += 1
	return i
}

func (cars Collection) GenerateCars(start, limit int) Collection {
	carChannel := make(chan *IndexedCar)
	var waitGroup sync.WaitGroup
	numCarsToGenerate := start + limit - 1
	generatedCars := Collection{}
	waitGroup.Add(numCarsToGenerate)
	next := carGenerator(iterator, start -1, numCarsToGenerate)
	carIndex, done := next()
	for !done {
		go func(carIndex int) {
			thisCar, err := GetThisCar(carIndex)
			if err != nil {
				panic(err)
			}
			carChannel <- thisCar
			waitGroup.Done()
		}(carIndex)

		carIndex, done = next()
	}

	go func() {
		waitGroup.Wait()
		println("close channel")
		close(carChannel)
	}()

	for thisCar := range carChannel {
		generatedCars = append(generatedCars, thisCar.Car)
	}
	return generatedCars
}

