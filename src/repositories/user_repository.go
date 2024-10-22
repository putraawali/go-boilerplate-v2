package repositories

import (
	"context"
	"errors"
	"fmt"
	"go-boilerplate-v2/src/constants"
	"go-boilerplate-v2/src/models"
	"go-boilerplate-v2/src/pkg/response"
	"net/http"

	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(ctx context.Context, user *models.User) (err error)
	FindByEmail(ctx context.Context, email string) (user models.User, err error)
}

type userRepository struct {
	db       *gorm.DB
	response *response.Response
}

func NewUserRepository(di di.Container) UserRepository {
	return &userRepository{
		db:       di.Get(constants.PG_DB).(*gorm.DB),
		response: di.Get(constants.RESPONSE).(*response.Response),
	}
}

func (u *userRepository) Insert(ctx context.Context, user *models.User) (err error) {
	if err = u.db.Create(&user).WithContext(ctx).Error; err != nil {
		err = u.response.NewError().
			SetContext(ctx).
			SetDetail(err.Error()).
			SetMessage(err).
			SetStatusCode(http.StatusInternalServerError)
	}

	return
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (user models.User, err error) {
	if err = u.db.First(&user, "email = ?", email).WithContext(ctx).Error; err != nil {
		code := http.StatusInternalServerError
		msg := err

		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
			msg = fmt.Errorf("user with email %s not found", email)
		}

		err = u.response.NewError().
			SetContext(ctx).
			SetDetail(msg.Error()).
			SetMessage(msg).
			SetStatusCode(code)
	}

	return
}
