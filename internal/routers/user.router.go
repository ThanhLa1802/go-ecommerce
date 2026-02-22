package routers

import (
	"go-ecommerce-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//public router
	// userRouterPublic := Router.Group("/product")
	// {
	// 	userRouterPublic.POST("/register")
	// 	userRouterPublic.POST("/otp")
	// }
	//private router
	userController, _ := wire.InitUserRouterHandler()
	userRouterPrivate := Router.Group("/user")

	{
		userRouterPrivate.GET("/get_user_by_id", userController.GetUserById)
	}

}
