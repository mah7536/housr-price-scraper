package domain

import (
	"encoding/json"
	"time"
)

func (t *TimeStamp) UnmarshalJSON(data []byte) (err error) {
	var timestamp int
	err = json.Unmarshal(data, &timestamp)
	if err != nil {
		return
	}

	tmp := TimeStamp(time.Unix(int64(timestamp), 0))

	*t = tmp
	return
}

func (t *TimeStamp) String() string {
	r := time.Time(*t)
	return r.Format(time.RFC3339)
}
