package main

import (
	"go.uber.org/fx"
	"todo-client/service"
	"todo-client/web"
	"todo-client/web/controllers"
	"todo-client/web/routes"
)

func main() {
	fx.New(
		fx.Provide(
			service.StartTaskAppClient,
		),
		fx.Invoke(
			controllers.StartController,
			routes.RegisterRoutes,
			web.Start,
		),
	).Run()
}
