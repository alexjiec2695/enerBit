package usecases

import (
	"github.com/google/wire"
	"rest_app/domain/usecase/enerbit"
)

// ProviderSet Provider of Injection Dependencies
var ProviderSet = wire.NewSet(
	enerbit.NewUseCase,
)
