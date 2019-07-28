package dss

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Zone contains information of a zone in the DSS.
type Zone struct {
	ZoneID float64
	name   string
}

// Takes a JSON result as []byte and converts it to Zone.
func createZoneFromJSON(data []byte) (Zone, error) {
	var jsonData map[string]interface{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return Zone{0, ""}, errors.New("Failed to unmarshal to zone -> " + err.Error())
	}
	return Zone{ZoneID: jsonData["ZoneID"].(float64), name: jsonData["name"].(string)}, nil
}

// String returns the zone as string.
func (z Zone) String() string {
	return fmt.Sprintf(
		"[%f : %s]",
		z.ZoneID,
		z.name)
}
