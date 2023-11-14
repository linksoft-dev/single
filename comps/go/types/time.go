package types

import (
	"encoding/xml"
	"time"
)

type Time struct {
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	} else {
		return t.Time.MarshalJSON()
	}
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	strData := string(data)
	if strData == "null" || strData == "0001-01-01T00:00:00Z" {
		return
	}
	// aspas dupla é necessario pq quando converte de bytes para string(string(data)) a string vem com aspas duplas
	formats := []string{
		"\"2006-01-02T15:04:05Z07:00\"",
		"\"2006-01-02T15:04:05Z\"",
		"\"2006-01-02 15:04:05\"",
	}
	var convertedTime time.Time
	for _, format := range formats {
		convertedTime, err = time.Parse(format, strData)
		if err == nil {
			break
		}
	}
	*t = Time{convertedTime}
	return
}

func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	formats := []string{
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05Z",
		"2006-01-02",
	}
	var v string
	var convertedTime time.Time
	if err = d.DecodeElement(&v, &start); err != nil {
		return
	}
	for _, format := range formats {
		convertedTime, err = time.Parse(format, v)
		if err == nil {
			break
		}
	}
	*t = Time{convertedTime}
	return
}
