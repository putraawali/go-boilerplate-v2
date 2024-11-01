package mock_usecases

import (
	"go-boilerplate-v2/src/usecases"

	"go.uber.org/mock/gomock"
)

func NewBaseUsecaseMock(ctrl *gomock.Controller) *usecases.Usecases {
	return &usecases.Usecases{
		User: NewMockUserUsecase(ctrl),
	}
}
