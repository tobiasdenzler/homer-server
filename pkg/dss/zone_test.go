package dss

import (
	"testing"
)

type testpair struct {
	input    []byte
	expected Zone
}

var tests = []testpair{
	{[]byte("{\"ZoneID\":1,\"name\":\"Wohnzimmer\"}"), Zone{ZoneID: 1, name: "Wohnzimmer"}},
	{[]byte("{\"ZoneID\":2,\"name\":\"K\u00FCche\"}"), Zone{ZoneID: 2, name: "Küche"}},
	{[]byte("{\"ZoneID\":29062,\"name\":\"keine Zuweisung\"}"), Zone{ZoneID: 29062, name: "keine Zuweisung"}},
	{[]byte("{\"ZoneID\":0,\"name\":\"\"}"), Zone{ZoneID: 0, name: ""}},
}

func TestCreateZoneFromJSON(t *testing.T) {

	for _, pair := range tests {
		v, err := createZoneFromJSON(pair.input)
		if err != nil {
			t.Errorf("Error during conversion of zone -> %s", err.Error())
		}
		if v != pair.expected {
			t.Errorf("Incorrect conversion of zone, got: %s, want: %s", v, pair.expected)
		}
	}
}
