package h

import (
	"testing"
	"time"

	"github.com/lib/pq"
)

func TestNullTimeCheck(t *testing.T) {
	test := pq.NullTime{
		Valid: true,
		Time:  time.Now(),
	}

	if NullTimeCheck(test) == "NULL" {
		t.Errorf("NullTimeCheck test failed")
	}
}

func TestInvalidTimeWillBeNull(t *testing.T) {
	if InvalidTimeWillBeNull("") != nil {
		t.Errorf("InvalidTimeWillBeNull test failed")
	}

	if InvalidTimeWillBeNull("fdsffds") != nil {
		t.Errorf("InvalidTimeWillBeNull test failed")
	}

	if InvalidTimeWillBeNull("01/31/2018") != nil {
		t.Errorf("InvalidTimeWillBeNull test failed")
	}

	if InvalidTimeWillBeNull("2018-01-31") != "2018-01-31" {
		t.Errorf("InvalidTimeWillBeNull test failed")
	}
}
