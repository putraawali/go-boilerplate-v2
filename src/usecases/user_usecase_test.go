package usecases_test

import (
	"go-boilerplate-v2/src/repositories"
	"go-boilerplate-v2/src/usecases"

	"github.com/stretchr/testify/suite"
)

type (
	userUsecaseTest struct {
		suite.Suite
		usecase usecases.UserUsecase
	}

	userUsecaseData struct {
		mockRepo repositories.Repositories
	}
)
