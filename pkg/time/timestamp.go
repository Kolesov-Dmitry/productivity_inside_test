package time

import (
	"strconv"
	"strings"
	"time"
)

// Timestamp is pseudonym for time.Time to use with JSON format as timestamp in milliseconds
type Timestamp int64

// Now returns current system Timestamp
func Now() Timestamp {
	return Timestamp(time.Now().UnixNano() / 1000)
}

// MarshalJSON marshals timestamp in JSON format
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(t), 10)), nil
}

// UnmarshalJSON unmarshalls timestamp JSON representation
func (t *Timestamp) UnmarshalJSON(s []byte) error {
	data := strings.Replace(string(s), "\"", "", -1)

	val, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return err
	}

	*(*int64)(t) = val

	return nil
}

// String returns timestamp as string
func (t Timestamp) String() string {
	return time.Unix(int64(t)/1000, 0).String()
}
