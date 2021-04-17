package app

import (
	"github.com/gin-gonic/gin"
	"hydra-sso/controllers"
	"hydra-sso/hydra_client"
)
type App struct {
	port string
	engine *gin.Engine
}

func NewApp() *App {
	return &App{
		port: ":3000",
	}
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func (app *App) Init() {
	engine := gin.Default()
	engine.Use(CORSMiddleware())
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/health", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	hydra := hydra_client.NewHydraClient("http://localhost:4445", "http://localhost:4444")
	loginController := controllers.NewLoginController(hydra)
	consentController := controllers.NewConsentController(hydra)
	userController := controllers.NewUserController(hydra)

	engine.GET("/login", loginController.Login)
	engine.POST("/login", loginController.AuthUsernamePassword)
	engine.GET("/consent", consentController.GetConsent)
	engine.POST("/consent", consentController.AuthConsent)
	engine.GET("/userinfor", userController.GetUserInformation)
	app.engine = engine
}

func (app *App) Run() {
		app.engine.Run(app.port)
}
