//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"suzushin54/event-sourcing-with-go/app/infrastructure"

	"suzushin54/event-sourcing-with-go/app/controller"
	"suzushin54/event-sourcing-with-go/app/infrastructure/http"
	"suzushin54/event-sourcing-with-go/app/infrastructure/repository_impl"
	"suzushin54/event-sourcing-with-go/app/usecase"
	"suzushin54/event-sourcing-with-go/pkg"
)

func InitServer(db *sqlx.DB, snowflake pkg.SnowflakeNode) *http.Server {
	wire.Build(
		controller.NewOrderController,
		usecase.NewOrderUseCaseV1,
		usecase.NewOrderUseCaseV2,
		repository_impl.NewOrderRepositoryV1,
		repository_impl.NewOrderRepositoryV2,
		infrastructure.NewMySqlEventStore,
		http.NewMux,
		http.NewServer,
	)
	return nil
}
