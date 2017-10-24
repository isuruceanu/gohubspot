package gohubspot

import "time"
import "strconv"
import "fmt"

type UnixTime struct {
	time.Time
}

var nilTime = (time.Time{}).UnixNano()

func (t *UnixTime) MarshalJSON() ([]byte, error) {
	nanos := t.UnixNano()
	if nanos == nilTime {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("\"%d\"", nanos/1000000)), nil
}

func (t *UnixTime) UnmarshalJSON(data []byte) error {
	millis, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	t.Time = time.Unix(0, millis*int64(time.Millisecond))

	return nil
}
