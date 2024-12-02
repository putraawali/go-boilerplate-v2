package usecases_test

import (
	"context"
	"errors"
	"go-boilerplate-v2/src/constants"
	"go-boilerplate-v2/src/dtos"
	src_mock "go-boilerplate-v2/src/mocks"
	"go-boilerplate-v2/src/models"
	"go-boilerplate-v2/src/repositories"
	mock_repositories "go-boilerplate-v2/src/repositories/mocks"
	"go-boilerplate-v2/src/usecases"
	"testing"

	"github.com/sarulabs/di"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type (
	userUsecaseTest struct {
		suite.Suite
		usecase usecases.UserUsecase
	}

	userUsecaseData struct {
		mockRepo *repositories.Repositories
		builder  *di.Builder
	}
)

func TestUserUsecase(t *testing.T) {
	suite.Run(t, new(userUsecaseTest))
}

func setupUserUsecase(t *testing.T) userUsecaseData {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repositories.NewMockRepository(ctrl)

	mockDI := src_mock.Dependencies{
		Repository: mockRepo,
	}

	return userUsecaseData{
		builder:  src_mock.NewMockDependencies(mockDI),
		mockRepo: mockRepo,
	}
}

func (u *userUsecaseTest) TestRegister() {
	ctx := context.WithValue(context.TODO(), constants.RequestID, "213")
	setup := setupUserUsecase(u.T())

	u.usecase = usecases.NewUserUsecase(setup.builder.Build())

	userRepo := setup.mockRepo.User.(*mock_repositories.MockUserRepository)

	param := dtos.RegisterParam{
		Email:     "user@mail.com",
		Password:  "Password1234",
		FirstName: "Dummy",
		LastName:  "User",
		Phone:     "081213123123",
	}

	u.Run("Success: success register", func() {
		userRepo.EXPECT().FindByEmail(ctx, param.Email).Return(models.User{}, nil)

		data := models.User{}
		data.FillRegister(param)

		userRepo.EXPECT().Insert(ctx, &data).Return(nil)

		err := u.usecase.Register(ctx, param)
		u.Nil(err, "err should be nil")
	})

	u.Run("Failed: error register new user", func() {
		userRepo.EXPECT().FindByEmail(ctx, param.Email).Return(models.User{}, nil)

		data := models.User{}
		data.FillRegister(param)

		userRepo.EXPECT().Insert(ctx, &data).Return(errors.New("mock error"))

		err := u.usecase.Register(ctx, param)
		u.NotNil(err, "err should be not nil")
	})

	u.Run("Failed: email already used", func() {
		userRepo.EXPECT().FindByEmail(ctx, param.Email).Return(models.User{
			UserID: 1,
		}, nil)

		err := u.usecase.Register(ctx, param)
		u.NotNil(err, "err should be not nil")
	})

	u.Run("Failed: failed find email", func() {
		userRepo.EXPECT().FindByEmail(ctx, param.Email).Return(models.User{}, errors.New("mock error"))

		err := u.usecase.Register(ctx, param)
		u.NotNil(err, "err should be not nil")
	})
}
