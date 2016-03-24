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
	Time    int    `json:"time"`
	Summary string `json:"summary"`
	Icon    string `json:"icon"`
	commonData
}

type minutelyData struct {
	Summary string       `json:"summary"`
	Icon    string       `json:"icon"`
	Data    []commonData `json:"data"`
}

type hourlyData struct {
	Summary string       `json:"summary"`
	Icon    string       `json:"icon"`
	Data    []commonData `json:"data"`
}

type dailyData struct {
	Summary string       `json:"summary"`
	Icon    string       `json:"icon"`
	Data    []commonData `json:"data"`
}

type commonData struct {
	ApparentTemperature        float64 `json:"apparentTemperature"`
	ApparentTemperatureMax     float64 `json:"apparentTemperatureMax"`
	ApparentTemperatureMaxTime int     `json:"apparentTemperatureMaxTime"`
	ApparentTemperatureMin     float64 `json:"apparentTemperatureMin"`
	ApparentTemperatureMinTime int     `json:"apparentTemperatureMinTime"`
	CloudCover                 float64 `json:"cloudCover"`
	DewPoint                   float64 `json:"dewPoint"`
	Humidity                   float64 `json:"humidity"`
	Icon                       string  `json:"icon"`
	MoonPhase                  float64 `json:"moonPhase"`
	NearestStormBearing        int     `json:"nearestStormBearing"`
	NearestStormDistance       int     `json:"nearestStormDistance"`
	Ozone                      float64 `json:"ozone"`
	PrecipIntensity            float64 `json:"precipIntensity"`
	PrecipIntensityMax         float64 `json:"precipIntensityMax"`
	PrecipProbability          float64 `json:"precipProbability"`
	Pressure                   float64 `json:"pressure"`
	Summary                    string  `json:"summary"`
	SunriseTime                int     `json:"sunriseTime"`
	SunsetTime                 int     `json:"sunsetTime"`
	Temperature                float64 `json:"temperature"`
	TemperatureMax             float64 `json:"temperatureMax"`
	TemperatureMaxTime         int     `json:"temperatureMaxTime"`
	TemperatureMin             float64 `json:"temperatureMin"`
	TemperatureMinTime         int     `json:"temperatureMinTime"`
	Time                       int     `json:"time"`
	Visibility                 float64 `json:"visibility"`
	WindBearing                int     `json:"windBearing"`
	WindSpeed                  float64 `json:"windSpeed"`
}

type flagsData struct {
	Sources         []string `json:"sources"`
	DarkskyStations []string `json:"darksky-stations"`
	LampStations    []string `json:"lamp-stations"`
	IsdStations     []string `json:"isd-stations"`
	MadisStations   []string `json:"madis-stations"`
	Units           string   `json:"units"`
}
