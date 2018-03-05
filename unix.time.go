package gohubspot

import "time"
import "strconv"
import "fmt"

type UnixTime struct {
	time.Time
}

var nilTime = (time.Time{}).UnixNano()

func (t *UnixTime) ToDate() {
	nanos := t.UnixNano()
	if nanos == nilTime {
		return
	}

	t.Time = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func (t *UnixTime) String() string {
	nanos := t.UnixNano()
	if nanos == nilTime {
		return ""
	}
	return fmt.Sprintf("%d", nanos/1000000)
}

func (t *UnixTime) MarshalJSON() ([]byte, error) {
	str := t.String()

	if len(str) == 0 {
		return []byte("null"), nil
	}

	return []byte(str), nil
}

func (t *UnixTime) UnmarshalJSON(data []byte) error {
	millis, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	t.Time = time.Unix(0, millis*int64(time.Millisecond))

	return nil
}
