package types

type WeatherCondition struct {
	ID          string
	Temperature float64
	WindSpeed   float64
}

type Service interface {
	CreateWeatherCondition(temperature float64, windSpeed float64) (*WeatherCondition, error)
	DeleteWeatherCondition(id string) error
	UpdateWeatherCondition(id string, temperature *float64, windSpeed *float64) error
	ReadWeatherCondition(id string) (*WeatherCondition, error)
	ListWeatherConditions() ([]*WeatherCondition, error)
}
