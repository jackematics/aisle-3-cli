package isodate

import (
	"encoding/json"
	"time"
)

type Frequency struct {
	Weeks int `json:"weeks,omitempty"`
	Months int `json:"months,omitempty"`
	Years int `json:"years,omitempty"`
}

type ISODate time.Time

func (date *ISODate) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = str[1:len(str)-1]

	parsed, err := time.Parse("2006-01-02", str)

	if err != nil {
		return err
	}

	*date = ISODate(parsed)
	return nil
}

func (date ISODate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(date).Format("2006-01-02"))
}

func (date ISODate) ToISOString() string {
	return time.Time(date).Format("2006-01-02")
}
