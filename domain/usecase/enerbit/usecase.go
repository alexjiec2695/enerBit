package enerbit

import (
	"rest_app/domain/model/entities/gateways"
)

type UseCase struct {
	gateway gateways.Gateway
}

func NewUseCase(gateway gateways.Gateway) *UseCase {
	return &UseCase{gateway}
}

func (usecase *UseCase) TestUsecase() string {
	return usecase.gateway.GetTest()
}
