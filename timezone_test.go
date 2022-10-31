package timezone

import (
    "testing"
    "regexp"
)

// TestGetTimezone calls timezone.GetTimezone with region & city, checking
// for a valid return value. Supports daylight and rest of year.
func TestGetUSTimezone(t *testing.T) {
    msg, err := GetTimezone("America", "Chicago")
    if (!regexp.MustCompile("CST").MatchString(msg) && !regexp.MustCompile("CDT").MatchString(msg)) || err != nil {
        t.Fatalf(`GetTimezone("America","Chicago") = %q, %v, want match for %#q, nil`, msg, err, "CST|CDT")
    }
}

// TestGetTimezone calls timezone.GetTimezone with region & city, checking
// for a valid return value.
func TestGetEUTimezone(t *testing.T) {
    want := regexp.MustCompile("CET")
    msg, err := GetTimezone("Europe", "Stockholm")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`GetTimezone("Europe","Sweden") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestGetTimezone calls timezone.GetTimezone with country & city, checking
// for an error.
func TestGetInvalidTimezone(t *testing.T) {
    msg, err := GetTimezone("Sweden", "Stockholm")
    if msg != "" || err == nil {
        t.Fatalf(`GetTimezone("Sweden", "Stockholm") = %q, %v, want "", error`, msg, err)
    }
}
