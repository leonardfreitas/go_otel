package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"leonardfreitas/go_otel/src/service-b/internal/dto"
	"leonardfreitas/go_otel/src/service-b/internal/entity"
	"leonardfreitas/go_otel/src/service-b/pkg/adapter/errorhandle"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type controller struct {
	usecase   entity.WeatherUseCase
	validator *validator.Validate
}

func NewWebController(usecase entity.WeatherUseCase, validator *validator.Validate) entity.WeatherController {
	return &controller{
		usecase:   usecase,
		validator: validator,
	}
}

var ErrNotFound = errors.New("not found")

func (controller controller) Get(response http.ResponseWriter, request *http.Request) {
	carrier := propagation.HeaderCarrier(request.Header)
	ctx := request.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	time.Sleep(time.Millisecond * 1000)

	cep, err := dto.FromQueryStringRequestToCep(chi.URLParam(request, "cep"), controller.validator)

	if err != nil {
		statusCode, message := errorhandle.Handle(errorhandle.ErrUnprocessableEntity)
		response.WriteHeader(statusCode)
		json.NewEncoder(response).Encode(message)

		return
	}

	weather, err := controller.usecase.Get(ctx, cep.Cep)

	if err != nil {
		response.WriteHeader(404)
		json.NewEncoder(response).Encode(map[string]string{"error": "can not find zipcode"})
		return
	}

	json.NewEncoder(response).Encode(weather)
}
