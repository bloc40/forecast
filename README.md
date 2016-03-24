# forecast
Go API wrapper for the [forecast.io](https://developer.forecast.io/docs/v2) weather service.

## How to use
Make sure to replace API_KEY string with your own key.

```go
package main

import (
	"fmt"
	"log"

	"github.com/bloc40/forecast"
)

func main() {
  data, err := forecast.Get("API_KEY", 37.3541667, -121.9541667)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(data)
  fmt.Println("Timezone:", data.Timezone)
  fmt.Println("Currently summary:", data.Currently.Summary)
  fmt.Println("Currently icon:", data.Currently.Icon)

  fmt.Println("Daily forecast summary:", data.Daily.Summary)
  fmt.Println("Daily forecast icon:", data.Daily.Icon)

  for _, dailyData := range data.Daily.Data {
    fmt.Println("Time:", dailyData.Time)
    fmt.Println("Summary:", dailyData.Summary)
    fmt.Println("Icon:", dailyData.Icon)
    fmt.Println("SunriseTime:", dailyData.SunriseTime)
    fmt.Println("SunsetTime:", dailyData.SunsetTime)
  }
}
```
