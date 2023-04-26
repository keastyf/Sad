package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestTrimNameFunction(t *testing.T) {
	is := is.New(t)

	_, err := trimName("/Enterprise Server Mitthem/IoT-gränssnitt/MQTT-klient/!UC_Framåt/UC_FRAMÅT_VV1_EM01-T2/Value")
	is.NoErr(err)
}

func TestTransformDataFunction(t *testing.T) {
	is := is.New(t)

	gurka := transformData([]byte(kyckling))
	is.Equal(len(gurka), 9)
}

const kyckling string = `[{
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
