package routes

import(
	handler "berbagi/controllers"
	//"berbagi/middlewares"
	//"fmt"
	//"reflect"
)

func registerRoutes() {
	e.POST("/register/donors", handler.RegisterDonorController)
}

