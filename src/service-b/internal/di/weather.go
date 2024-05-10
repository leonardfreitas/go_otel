package di

import (
	"net/http"

	"leonardfreitas/go_otel/src/service-b/internal/entity"
	"leonardfreitas/go_otel/src/service-b/internal/infra/viacep"
	"leonardfreitas/go_otel/src/service-b/internal/infra/weatherapi"
	"leonardfreitas/go_otel/src/service-b/internal/infra/web"
	"leonardfreitas/go_otel/src/service-b/internal/usecase"

	"github.com/go-playground/validator/v10"
)

func ConfigWebController(validator *validator.Validate) entity.WeatherController {
	httpClient := http.DefaultClient

	cepHttpClient := viacep.NewCepHTTPClient(httpClient)
	weatherHttpClient := weatherapi.NewWeatherHTTPClient(httpClient)
	weatherUseCase := usecase.NewWeatherUseCase(cepHttpClient, weatherHttpClient)
	weatherController := web.NewWebController(weatherUseCase, validator)

	return weatherController
}
