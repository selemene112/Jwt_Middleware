package router

import (
	"jwt_token/controllers"
	"jwt_token/middlewares"

	"github.com/gin-gonic/gin"
)



func StartApp() *gin.Engine{
	r := gin.Default()
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)

		



	}

	productRoute := r.Group("/products")
	{
		productRoute.Use(middlewares.Authentication())
		productRoute.POST("/", controllers.CreateProduct)
		productRoute.GET("/:productId", middlewares.ProductAuthorization(),controllers.UserIndex)
		

		productRoute.PUT("/:productId", middlewares.ProductAuthorization(),controllers.UpdateProduct)

	}


	adminRoute := r.Group("/admin")
	{
		adminRoute.POST("/register", controllers.AdminRegister)
		adminRoute.POST("/login", controllers.AdminLogin)
		// adminRoute.GET("/Product", controllers.ShowProduct)

	}

	adminProductRoute := r.Group("/Aproducts")
	{
		adminProductRoute.Use(middlewares.Authentication())
		adminProductRoute.GET("/Product", controllers.AdminIndex)
		adminProductRoute.PUT("/:productId",controllers.UpdateProduct1)
		adminProductRoute.DELETE("/Product",controllers.Delete)
	}
	return r
}