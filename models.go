package forecast

type Data struct {
	Latitude  float64       `json:"latitude"`
	Longitude float64       `json:"longitude"`
	Timezone  string        `json:"timezone"`
	Offset    int           `json:"offset"`
	Currently currentlyData `json:"currently"`
	Minutely  minutelyData  `json:"minutely"`
	Hourly    hourlyData    `json:"hourly"`
	Daily     dailyData     `json:"daily"`
	Flags     flagsData     `json:"flags"`
}

type currentlyData struct {
	Time                 int     `json:"time"`
	Summary              string  `json:"summary"`
	Icon                 string  `json:"icon"`
	NearestStormDistance int     `json:"nearestStormDistance"`
	NearestStormBearing  int     `json:"nearestStormBearing"`
	PrecipIntensity      int     `json:"precipIntensity"`
	PrecipProbability    int     `json:"precipProbability"`
	Temperature          float64 `json:"temperature"`
	ApparentTemperature  float64 `json:"apparentTemperature"`
	DewPoint             float64 `json:"dewPoint"`
	Humidity             float64 `json:"humidity"`
	WindSpeed            float64 `json:"windSpeed"`
	WindBearing          int     `json:"windBearing"`
	Visibility           float64 `json:"visibility"`
	CloudCover           float64 `json:"cloudCover"`
	Pressure             float64 `json:"pressure"`
	Ozone                float64 `json:"ozone"`
}

type minutelyData struct {
	Summary string       `json:"summary"`
	Icon    string       `json:"icon"`
	Data    []minuteData `json:"data"`
}

type minuteData struct {
	Time              int `json:"time"`
	PrecipIntensity   int `json:"precipIntensity"`
	PrecipProbability int `json:"precipProbability"`
}

type hourlyData struct {
	Summary string     `json:"summary"`
	Icon    string     `json:"icon"`
	Data    []hourData `json:"data"`
}

type hourData struct {
	Time                int     `json:"time"`
	Summary             string  `json:"summary"`
	Icon                string  `json:"icon"`
	Temperature         float64 `json:"temperature"`
	DewPoint            float64 `json:"dewPoint"`
	Humidity            float64 `json:"humidity"`
	WindSpeed           float64 `json:"windSpeed"`
	PrecipIntensity     int     `json:"precipIntensity"`
	PrecipProbability   int     `json:"precipProbability"`
	ApparentTemperature float64 `json:"apparentTemperature"`
	WindBearing         int     `json:"windBearing"`
	Visibility          float64 `json:"visibility"`
	CloudCover          float64 `json:"cloudCover"`
	Pressure            float64 `json:"pressure"`
	Ozone               float64 `json:"ozone"`
}

type dailyData struct {
	Summary string    `json:"summary"`
	Icon    string    `json:"icon"`
	Data    []dayData `json:"data"`
}

type dayData struct {
	Time                       int     `json:"time"`
	Summary                    string  `json:"summary"`
	Icon                       string  `json:"icon"`
	SunriseTime                int     `json:"sunsetTime"`
	SunsetTime                 int     `json:"sunsetTime"`
	MoonPhase                  float64 `json:"moonPhase"`
	PrecipIntensity            float64 `json:"precipIntensity"`
	PrecipIntensityMax         float64 `json:"precipIntensityMax"`
	PrecipProbability          float64 `json:"precipProbability"`
	TemperatureMin             float64 `json:"temperatureMin"`
	TemperatureMinTime         int     `json:"temperatureMinTime"`
	TemperatureMax             float64 `json:"temperatureMax"`
	TemperatureMaxTime         int     `json:"temperatureMaxTime"`
	ApparentTemperatureMin     float64 `json:"apparentTemperatureMin"`
	ApparentTemperatureMinTime int     `json:"apparentTemperatureMinTime"`
	ApparentTemperatureMax     float64 `json:"apparentTemperatureMax"`
	ApparentTemperatureMaxTime int     `json:"apparentTemperatureMaxTime"`
	DewPoint                   float64 `json:"dewPoint"`
	Humidity                   float64 `json:"humidity"`
	WindSpeed                  float64 `json:"windSpeed"`
	WindBearing                int     `json:"windBearing"`
	Visibility                 float64 `json:"visibility"`
	CloudCover                 float64 `json:"cloudCover"`
	Pressure                   float64 `json:"pressure"`
	Ozone                      float64 `json:"ozone"`
}

type flagsData struct {
	Sources         []string `json:"sources"`
	DarkskyStations []string `json:"darksky-stations"`
	LampStations    []string `json:"lamp-stations"`
	IsdStations     []string `json:"isd-stations"`
	MadisStations   []string `json:"madis-stations"`
	Units           string   `json:"units"`
}
