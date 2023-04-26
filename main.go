package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type DTO struct {
	Name  string
	Value float64
	Unit  string
}

type Data struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

func main() {

	transformedData := transformData([]byte(sniderData))

	fmt.Println(transformedData)

	for _, object := range transformedData {
		if object.Unit == "°C" {
			fmt.Printf("temperature is: %f\n", object.Value) //TODO:change to lwm2m model
		} else if object.Unit == "Wh" || object.Unit == "W" {
			fmt.Printf("consumption is: %f\n", object.Value) //TODO:change to lwm2m model
		}
	}
}

func transformData(incomingData []byte) []DTO {
	data := []Data{}

	err := json.Unmarshal(incomingData, &data)
	if err != nil {
		panic(fmt.Errorf("it didn't work: %s", err.Error()))
	}

	transformedData := []DTO{}

	for _, object := range data {
		transformedObject := DTO{}

		trimmedName, err := trimName(object.Name)
		if err != nil {
			panic(err.Error())
		}

		transformedObject.Name = trimmedName

		parsedvalue, err := strconv.ParseFloat(object.Value, 64)
		if err != nil {
			panic(fmt.Errorf("couldn't parse float from string"))
		}

		transformedObject.Value = parsedvalue

		transformedObject.Unit = object.Unit

		transformedData = append(transformedData, transformedObject)
	}

	return transformedData
}

func trimName(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("can't paste empty name")
	}
	prefix := "/Enterprise Server Mitthem/IoT-gränssnitt/MQTT-klient/!UC_Framåt/"
	suffix := "/Value"

	trimmedprefix := strings.TrimPrefix(name, prefix)
	trimmedsuffix := strings.TrimSuffix(trimmedprefix, suffix)

	return trimmedsuffix, nil
}

const sniderData string = `[{
	"name":"/Enterprise Server Mitthem/IoT-gränssnitt/MQTT-klient/!UC_Framåt/UC_FRAMÅT_VV1_EM01-T2/Value",
	"value":"8",
	"unit":"°C",
	"description":"Returtemperatur Varmvatten"
	},{
	"name":"/Enterprise Server Mitthem/IoT-gränssnitt/MQTT-klient/!UC_Framåt/UC_FRAMÅT_OUTDOOR-TEMP/Value",
	"value":"2.1800000667572021",
	"unit":"°C",
	"description":"Utetemperatur"
	},{
	"name":"/Enterprise Server Mitthem/IoT-gränssnitt/MQTT-klient/!UC_Framåt/UC_FRAMÅT_VP1_EM01-ENERGY/Value",
	"value":"3372000000",
	"unit":"Wh",
	"description":"Mätarställning Värme Primär"
	},{
	"name":"/Enterprise Server Mitthem/IoT-gränssnitt/MQTT-klient/!UC_Framåt/UC_FRAMÅT_VP1_EM01-POWER/Value",
	"value":"66000",
	"unit":"W",
	"description":"Momentaneffekt Värme Primär"
	},{
	"name":"/Enterprise Server Mitthem/IoT-gränssnitt/MQTT-klient/!UC_Framåt/UC_FRAMÅT_VP1_EM01-T1/Value",
	"value":"76",
	"unit":"°C",
	"description":"Tilloppstemperatur Värme Primär"
	},{
	"name":"/Enterprise Server Mitthem/IoT-gränssnitt/MQTT-klient/!UC_Framåt/UC_FRAMÅT_VP1_EM01-T2/Value",
	"value":"25",
	"unit":"°C",
	"description":"Returtemperatur Värme Primär"
	},{
	"name":"/Enterprise Server Mitthem/IoT-gränssnitt/MQTT-klient/!UC_Framåt/UC_FRAMÅT_VV1_EM01-ENERGY/Value",
	"value":"215000000",
	"unit":"Wh",
	"description":"Mätarställning Varmvatten"
	},{
	"name":"/Enterprise Server Mitthem/IoT-gränssnitt/MQTT-klient/!UC_Framåt/UC_FRAMÅT_VV1_EM01-POWER/Value",
	"value":"12000",
	"unit":"W",
	"description":"Momentaneffekt Varmvatten"
	},{
	"name":"/Enterprise Server Mitthem/IoT-gränssnitt/MQTT-klient/!UC_Framåt/UC_FRAMÅT_VV1_EM01-T1/Value",
	"value":"53",
	"unit":"°C",
	"description":"Tilloppstemperatur Varmvatten"
	}]`
