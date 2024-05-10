package entity

import (
	"context"
	"net/http"

	"leonardfreitas/go_otel/src/service-a/internal/dto"
)

type WeatherHTTPClient interface {
	Get(context.Context, string) (*dto.WeatherOutput, error)
}

type WeatherUseCase interface {
	Get(context.Context, string) (*dto.WeatherOutput, error)
}

type WeatherController interface {
	Get(http.ResponseWriter, *http.Request)
}
