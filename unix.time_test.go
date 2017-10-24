package gohubspot

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

var epochStr = "1479298304451"

func TestMillisecondsEpochToDateTime(t *testing.T) {

	epoch := []byte(epochStr) // = 2016-11-16 14:11:44.451 +0200 EET
	result := &UnixTime{}

	if err := result.UnmarshalJSON(epoch); err != nil {
		t.Error(err)
	}

	if err := checkTime(result); err != nil {
		t.Error(err)
	}
}

func TestCanMarshalUnixTime(t *testing.T) {
	ti, err := time.Parse(time.RFC3339, "2016-11-16T14:11:44.451+02:00")
	if err != nil {
		t.Error(err)
	}
	unixTime := &UnixTime{Time: ti}

	bytes, err := unixTime.MarshalJSON()
	if err != nil {
		t.Error(err)
	}

	unixTimeStr := strings.Replace(string(bytes[:]), "\"", "", -1)

	if unixTimeStr != epochStr {
		t.Errorf("MarshalTime expected %s, but %s", epochStr, unixTimeStr)
	}
}

func TestCanUnmarshalUnixTimeJson(t *testing.T) {
	model := struct {
		Size int      `json:"size"`
		Time UnixTime `json:"created_at"`
	}{}
	jsonObj := `
		{"size": 34, "created_at":%s}
	`
	data := fmt.Sprintf(jsonObj, epochStr)

	if err := json.Unmarshal([]byte(data), &model); err != nil {
		t.Error(err)
	}

	if err := checkTime(&model.Time); err != nil {
		t.Error(err)
	}

}

func checkTime(unixTime *UnixTime) error {

	if unixTime.Year() != 2016 {
		return fmt.Errorf("Year expected %d, but %d", 2016, unixTime.Year())
	}

	if unixTime.Month() != 11 {
		fmt.Errorf("Month expected %d, but %d", 11, unixTime.Month())
	}

	if unixTime.Second() != 44 {
		fmt.Errorf("Seconds expected %d, but %d", 44, unixTime.Second())
	}

	return nil
}
