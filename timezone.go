package timezone
import "errors"

// GetTimezone returns timezone for given country / city.
func GetTimezone(country string, city string) (string, error) {
  if country == "" && city == "" {
    return "", errors.New("empty country or city")
  }

  message := "Not implemented"
  return message, nil
}
