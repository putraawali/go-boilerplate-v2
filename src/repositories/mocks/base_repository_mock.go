package mock_repositories

import (
	"go-boilerplate-v2/src/repositories"

	"go.uber.org/mock/gomock"
)

func NewMockRepository(ctrl *gomock.Controller) *repositories.Repositories {
	return &repositories.Repositories{
		User: NewMockUserRepository(ctrl),
	}
}
