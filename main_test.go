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
