package timezone
import "errors"
import "time"
import "fmt"

func GetTimezone(country string, city string) (string, error) {                 // GetTimezone returns timezone for given city
  if country == "" && city == "" {
    return "", errors.New("empty country or city")
  }

  loc, err := time.LoadLocation(country+`/`+city)                                 // "Europe/Stockholm"
  if err != nil {
    fmt.Println(err)
    return "", err
  }

  t := time.Now().In(loc)

  return t.Format("MST"), nil                                                   // Return timezone
}
