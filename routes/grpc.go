package routes

import (
	proto "github.com/goravel/example-proto"

	"dakhl/app/facades"
	"dakhl/app/grpc/controllers"
	httpcontrollers "dakhl/app/http/controllers"
)

func Grpc() {
	proto.RegisterUserServiceServer(facades.Grpc().Server(), controllers.NewUserController())

	grpcController := httpcontrollers.NewGrpcController()
	facades.Route().Get("/grpc/user", grpcController.User)
}
