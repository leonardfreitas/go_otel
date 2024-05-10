package di

import (
	"net/http"

	"leonardfreitas/go_otel/src/service-a/internal/entity"
	"leonardfreitas/go_otel/src/service-a/internal/infra/weatherapi"
	"leonardfreitas/go_otel/src/service-a/internal/infra/web"
	"leonardfreitas/go_otel/src/service-a/internal/usecase"

	"github.com/go-playground/validator/v10"
)

func ConfigWebController(validator *validator.Validate) entity.WeatherController {
	httpClient := http.DefaultClient

	weatherHttpClient := weatherapi.NewWeatherHTTPClient(httpClient)
	weatherUseCase := usecase.NewWeatherUseCase(weatherHttpClient)
	weatherController := web.NewWebController(weatherUseCase, validator)

	return weatherController
}
