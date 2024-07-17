package date

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// Value returns a driver Value.
// Value must not panic.
func (d Date) Value() (driver.Value, error) {
	return d.String(), nil
}

// Scan assigns a value from a database driver.
// The src value will be of one of the following types:
//
//	int64
//	float64
//	bool
//	[]byte
//	string
//	time.Time
//	nil - for NULL values
//
// An error should be returned if the value cannot be stored
// without loss of information.
//
// Reference types such as []byte are only valid until the next call to Scan
// and should not be retained. Their underlying memory is owned by the driver.
// If retention is necessary, copy their values before the next call to Scan.
func (d *Date) Scan(src any) error {
	switch v := src.(type) {
	case string:
		newVal, err := Parse(v)
		if err != nil {
			return fmt.Errorf("failed to parse string %+v to date", v)
		}
		*d = newVal
	case time.Time:
		d.t = v
	case []byte:
		newVal, err := Parse(string(v))
		if err != nil {
			return fmt.Errorf("failed to parse string %+v to date", v)
		}
		*d = newVal
	default:
		return fmt.Errorf("unexpected date type %+t", v)
	}
	return nil
}
