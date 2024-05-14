package web

import (
	"encoding/json"
	"net/http"
	"time"

	"leonardfreitas/go_otel/src/service-a/internal/dto"
	"leonardfreitas/go_otel/src/service-a/internal/entity"
	"leonardfreitas/go_otel/src/service-a/pkg/adapter/errorhandle"

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

func (controller controller) Get(response http.ResponseWriter, request *http.Request) {
	print("SERVICO A")
	carrier := propagation.HeaderCarrier(request.Header)
	ctx := request.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	tr := otel.Tracer("microservice-trace")
	ctx, span := tr.Start(ctx, "get weather")
	defer span.End()

	time.Sleep(time.Millisecond * 1000)

	cep, err := dto.FromQueryStringRequestToCep(chi.URLParam(request, "cep"), controller.validator)

	if err != nil {
		statusCode, message := errorhandle.Handle(errorhandle.ErrUnprocessableEntity)
		response.WriteHeader(statusCode)
		json.NewEncoder(response).Encode(message)

		return
	}

	weatherOutput, err := controller.usecase.Get(ctx, cep.Cep)

	if err != nil {
		print("ENTROU AQUIIIII")
		response.WriteHeader(404)
		json.NewEncoder(response).Encode(map[string]string{"error": "can not find zipcode"})
		return
	}

	json.NewEncoder(response).Encode(weatherOutput)
}
