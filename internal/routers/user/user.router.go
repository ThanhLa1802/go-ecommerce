package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *ProductRouter) InitUserRouter(Router *gin.RouterGroup) {
	//public router
	userRouterPublic := Router.Group("/product")
	{
		userRouterPublic.POST("/register")
		userRouterPublic.POST("/otp")
	}
	//private router
	userRouterPrivate := Router.Group("/user")

	{
		userRouterPrivate.POST("/active_user")
	}

}
