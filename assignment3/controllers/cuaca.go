package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type cuaca struct {
	Water       int
	Wind        int
	StatusWater string
	ClassWater  string
	StatusWind  string
	ClassWind   string
}

type datajson struct {
	Water int
	Wind  int
}

var filename = "dataJSON/cuaca.json"
var filenamedata = "dataJSON/datajson.json"

func cekCuaca() []cuaca {

	water, wind := genarateCuaca()

	var statusWater string
	var classWater string
	var statusWind string
	var classWind string

	if water <= 5 {
		statusWater = "AMAN"
		classWater = "bg-success"
	} else if water >= 6 && water <= 8 {
		statusWater = "SIAGA"
		classWater = "bg-warning"
	} else {
		statusWater = "BAHAYA"
		classWater = "bg-danger"
	}

	if wind <= 6 {
		statusWind = "AMAN"
		classWind = "bg-success"
	} else if wind >= 7 && wind <= 15 {
		statusWind = "SIAGA"
		classWind = "bg-warning"
	} else {
		statusWind = "BAHAYA"
		classWind = "bg-danger"
	}

	setData := cuaca{
		Water:       water,
		Wind:        wind,
		StatusWater: statusWater,
		ClassWater:  classWater,
		StatusWind:  statusWind,
		ClassWind:   classWind,
	}

	setDataJson := datajson{
		Water: water,
		Wind:  wind,
	}

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {

		dataCuaca := []cuaca{}

		dataCuaca = append(dataCuaca, setData)
		jsonString, _ := json.Marshal(dataCuaca)
		ioutil.WriteFile(filename, jsonString, os.ModePerm)
		jsonString2, _ := json.Marshal(setDataJson)
		ioutil.WriteFile(filenamedata, jsonString2, os.ModePerm)
		return dataCuaca
	}

	plan, _ := ioutil.ReadFile(filename)
	dataJSON := []cuaca{}
	err := json.Unmarshal(plan, &dataJSON)
	if err != nil {
		fmt.Printf(err.Error())
	}
	dataCuaca := []cuaca{}
	dataCuaca = append(dataCuaca, setData)

	dataJSON = append(dataCuaca, dataJSON...)
	jsonString, _ := json.Marshal(dataJSON)

	ioutil.WriteFile(filename, jsonString, os.ModePerm)

	jsonString2, _ := json.Marshal(setDataJson)
	ioutil.WriteFile(filenamedata, jsonString2, os.ModePerm)

	return dataJSON

}

func genarateCuaca() (water, wind int) {
	rand.Seed(time.Now().UnixNano())
	setWater := rand.Intn(100)
	setWind := rand.Intn(100)

	return setWater, setWind
}

func GetCuacaHandler(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("platform")

	if userAgent == "WEB" || userAgent == "" {

		getcuaca := cekCuaca()

		// process as html
		tpl, err := template.ParseFiles("views/static/index.html", "views/static/header.html", "views/static/note.html")
		if err != nil {
			writeJsonResponse(w, http.StatusNotFound, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		tpl.Execute(w, getcuaca)
		return
	}
}

func writeJsonResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
