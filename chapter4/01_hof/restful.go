package hof

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	"log"
	"strconv"
)

func CarsIndexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	response, err := getAllCarsJson();
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w	, string(response))
}

func CarHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	carIndex, err := strconv.Atoi(p[0].Value)
	if err != nil {
		log.Fatal("CarHandler unable to find car (%v) by index\n", p[0].Value)
	}
	response, err := getThisCarJson(carIndex);
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w	, string(response))
}


func getAllCarsJson() ([]byte, error) {
	return json.MarshalIndent(Payload{CarsDB}, "", "  ")
}

func getThisCarJson(carIndex int) ([]byte, error) {
	return json.MarshalIndent(CarsDB[carIndex], "", "  ")
}

func GetThisCar(carIndex int) (*IndexedCar, error) {
	thisCarJson, err  := getThisCarJson(carIndex)
	if err != nil {
		panic(err)
	}
	var thisIndexedCar IndexedCar
	if err := json.Unmarshal(thisCarJson, &thisIndexedCar); err != nil {
		panic(err)
	}
	return &thisIndexedCar, err
}
