package routes

import (
	"github.com/BRIZINGR007/app-002-code-assistant/internal/controllers"
	"github.com/BRIZINGR007/go-service-utils/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(server *gin.Engine) {
	baseServer := server.Group("/code-assistant")
	baseServer.GET("/healthcheck", func(context *gin.Context) {
		context.JSON(200, gin.H{"status": "OK"})
	})
	//User Routes
	userRoutes(baseServer)
}

func userRoutes(baseServer *gin.RouterGroup) {
	users_base_path := "/users"

	//Un-Authenticated Routes  ...
	user_ua := baseServer.Group(users_base_path)
	user_ua.POST("/signup", controllers.SignUp)
	user_ua.POST("/login", controllers.LogIn)
	user_ua.POST("/logout", controllers.LogOut)

	user := baseServer.Group(users_base_path)
	user.Use(middlewares.RestMiddleware)
	user.GET("/validate-session", controllers.ValidateSession)
	user.GET("/get-code-bases", controllers.GetAllCodeBases)

	//Authenticated Routes  ...
	codeassist_base_path := "/code-assist"
	codeassist := baseServer.Group(codeassist_base_path)
	codeassist.Use(middlewares.RestMiddleware)
	codeassist.POST("/extract-code", controllers.ExtractCode)
	codeassist.GET("/code-base-chat", controllers.CodeBaseChat)

	//AuthenticatedRoutes
	chat_base_path := "/chat"
	chat := baseServer.Group(chat_base_path)
	chat.Use(middlewares.RestMiddleware)
	chat.GET("get-all-chats", controllers.GetAllChats)

}
