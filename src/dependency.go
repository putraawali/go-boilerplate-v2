package src

import (
	"go-boilerplate-v2/src/constants"
	"go-boilerplate-v2/src/pkg/connections"
	"go-boilerplate-v2/src/pkg/response"
	"go-boilerplate-v2/src/repositories"
	"go-boilerplate-v2/src/usecases"

	"github.com/sarulabs/di"
)

func dependencyInjection() di.Container {
	builder, _ := di.NewBuilder()

	pg, err := connections.NewPostgreConnection()
	// mysql, err := connections.NewMySQLConnection()

	builder.Add(
		di.Def{
			Name: constants.RESPONSE,
			Build: func(ctn di.Container) (interface{}, error) {
				return response.NewResponse(), nil
			},
		},
		// di.Def{
		// 	Name: constants.MYSQL_DB,
		// 	Build: func(ctn di.Container) (interface{}, error) {
		// 		return mysql, err
		// 	},
		// },
		di.Def{
			Name: constants.PG_DB,
			Build: func(ctn di.Container) (interface{}, error) {
				return pg, err
			},
		},
		di.Def{
			Name: constants.REPOSITORY,
			Build: func(ctn di.Container) (interface{}, error) {
				return repositories.NewRepository(builder.Build()), nil
			},
		},
		di.Def{
			Name: constants.USECASE,
			Build: func(ctn di.Container) (interface{}, error) {
				return usecases.NewUsecase(builder.Build()), nil
			},
		},
	)

	return builder.Build()
}
