package timezone

import (
    "testing"
    "regexp"
)

// TestGetTimezone calls timezone.GetTimezone with country & city, checking
// for a valid return value.
func TestGetTimezone(t *testing.T) {
    want := regexp.MustCompile("CST")
    msg, err := GetTimezone("America", "Chicago")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`GetTimezone("America","Chicago") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}
