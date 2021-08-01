package users

import "github.com/kashyaprahul94/go-boilerplate/internal/http"

func InitApiRoutes(router http.HttpRouter) {
	userRouter := router.Group("/users")

	userRouter.GET("", GetUsers)
	userRouter.GET("/", GetUsers)

	userRouter.GET("/:userId", GetUserById)

	userRouter.POST("", CreateUser)
	userRouter.POST("/", CreateUser)

	userRouter.PUT("/:userId", UpsertUser)
	userRouter.PATCH("/:userId", PartialUpdateUser)

	userRouter.DELETE("/:userId", DeleteUser)
}
