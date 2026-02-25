package user

import (
	"go-ecommerce-backend-api/internal/controller"
	"go-ecommerce-backend-api/internal/repo"
	"go-ecommerce-backend-api/internal/service"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//public router
	ur := repo.NewUserRepo()
	us := service.NewUserService(ur)
	userHandlerNonDependency := controller.NewUserController(us)
	userRouterPublic := Router.Group("/product")
	{
		userRouterPublic.POST("/register")
		userRouterPublic.POST("/otp")
	}
	//private router
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("/get_user_by_id", userHandlerNonDependency.GetUserByID)
		userRouterPrivate.POST("/register", userHandlerNonDependency.RegisterUser)
	}

}
