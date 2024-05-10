package usecase

import (
	"context"

	"leonardfreitas/go_otel/src/service-a/internal/dto"
	"leonardfreitas/go_otel/src/service-a/internal/entity"
)

type usecase struct {
	weatherHTTPClient entity.WeatherHTTPClient
}

func NewWeatherUseCase(
	weatherHTTPClient entity.WeatherHTTPClient,
) entity.WeatherUseCase {
	return &usecase{
		weatherHTTPClient: weatherHTTPClient,
	}
}

func (usecase usecase) Get(ctx context.Context, cep string) (*dto.WeatherOutput, error) {
	weatherResponse, err := usecase.weatherHTTPClient.Get(ctx, cep)

	if err != nil {
		return nil, err
	}

	return weatherResponse, nil
}
