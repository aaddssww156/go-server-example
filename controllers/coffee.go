package controllers

import (
	"encoding/json"
	"net/http"
	"server-example/services"
	"server-example/tools"
)

var coffee services.Coffe

// GET/coffees
func GetAllCoffees(w http.ResponseWriter, r *http.Request) {
	all, err := coffee.GetAllCoffes()
	if err != nil {
		tools.MessageLogs.ErrorLog.Println(err)
		return
	}

	tools.WriteJson(w, http.StatusOK, tools.Envelope{"coffees": all})
}

func CreateCoffee(w http.ResponseWriter, r *http.Request) {
	var coffeeData services.Coffe
	if err := json.NewDecoder(r.Body).Decode(&coffeeData); err != nil {
		tools.MessageLogs.ErrorLog.Println(err)
		return
	}

	tools.WriteJson(w, http.StatusOK, coffeeData)

	coffeCreated, err := coffee.CreateCoffee(coffeeData)
	if err != nil {
		tools.MessageLogs.ErrorLog.Println(err)
	}

	tools.WriteJson(w, http.StatusOK, coffeCreated)
}
