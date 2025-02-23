package src

import (
	"go-boilerplate-v2/src/constants"
	"go-boilerplate-v2/src/pkg/connections"
	"go-boilerplate-v2/src/pkg/response"
	"go-boilerplate-v2/src/repositories"
	"go-boilerplate-v2/src/usecases"

	godi "github.com/putraawali/go-di"
)

func dependencyInjection() godi.Container {
	builder := godi.New()

	pg, err := connections.NewPostgreConnection()
	if err != nil {
		panic(err)
	}
	// mysql, err := connections.NewMySQLConnection()

	builder.Set(constants.RESPONSE, response.NewResponse())
	builder.Set(constants.PG_DB, pg)

	builder.Add(
		godi.Dependency{
			Name: constants.REPOSITORY,
			Create: func() (interface{}, error) {
				return repositories.NewRepository(builder.Build()), nil
			},
		},
		godi.Dependency{
			Name: constants.USECASE,
			Create: func() (interface{}, error) {
				return usecases.NewUsecase(builder.Build()), nil
			},
		},
	)

	return builder.Build()
}
