package src_mock

import (
	"go-boilerplate-v2/src/constants"
	"go-boilerplate-v2/src/pkg/response"
	"go-boilerplate-v2/src/repositories"

	godi "github.com/putraawali/go-di"
	"gorm.io/gorm"
)

type Dependencies struct {
	Repository *repositories.Repositories
	Postgres   *gorm.DB
	Mysql      *gorm.DB
}

func NewMockDependencies(d Dependencies) *godi.Builder {
	builder := godi.New()

	builder.Add(
		godi.Dependency{
			Name: constants.RESPONSE,
			Create: func() (interface{}, error) {
				return response.NewResponse(), nil
			},
		},
		godi.Dependency{
			Name: constants.REPOSITORY,
			Create: func() (interface{}, error) {
				return d.Repository, nil
			},
		},
		godi.Dependency{
			Name: constants.MYSQL_DB,
			Create: func() (interface{}, error) {
				return d.Mysql, nil
			},
		},
		godi.Dependency{
			Name: constants.PG_DB,
			Create: func() (interface{}, error) {
				return d.Postgres, nil
			},
		},
	)

	return builder
}
