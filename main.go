package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	bacon := []Data{}

	err := json.Unmarshal([]byte(sniderData), &bacon)
	if err != nil {
		panic(fmt.Errorf("it didnt work: %s", err.Error()))
	}

	cookedbacon := []DTO{}

	for _, value := range bacon {
		cookedslice := DTO{}

		trimedName, err := trimName(value.Name)
		if err != nil {
			panic(err.Error())
		}

		cookedslice.Name = trimedName

		parsedvalue, err := strconv.ParseFloat(value.Value, 64)
		if err != nil {
			panic(fmt.Errorf("couldnt parse float from string"))
		}

		cookedslice.Value = parsedvalue

		cookedbacon = append(cookedbacon, cookedslice)
	}

	fmt.Println(cookedbacon)
}

func trimName(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("cant paste empty name")
	}
	prefix := "/Enterprise Server Mitthem/IoT-gränssnitt/MQTT-klient/!UC_Framåt/"
	suffix := "/Value"

	trimedprefix := strings.TrimPrefix(name, prefix)

	trimedsuffix := strings.TrimSuffix(trimedprefix, suffix)

	return trimedsuffix, nil
}

type DTO struct {
	Name  string
	Value float64
}

type Data struct {
	Name  string `json:"name"`
	Value string `json:"value"`
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
