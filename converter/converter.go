package converter

import "github.com/farshidtz/senml/v2"

const SensorValue string = "5700"

type Converter interface {
	Convert(name, unit string, value float64) senml.Pack
}

type converter struct {
	prefix map[string]string
}

func NewConverter() Converter {
	return &converter{
		prefix: map[string]string{
			"°C": "urn:oma:lwm2m:ext:3303",
			"W":  "urn:oma:lwm2m:ext:3328",
			"Wh": "urn:oma:lwm2m:ext:3331",
		},
	}
}

func (c *converter) Convert(name, unit string, value float64) senml.Pack {
	pack := senml.Pack{}

	pack = append(pack, senml.Record{
		BaseName: c.prefix[unit],
		Name:     "0",
		Value:    &value,
	})

	return pack
}
