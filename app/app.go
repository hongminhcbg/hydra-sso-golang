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

func (app *App) Init() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/health", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	hydra := hydra_client.NewHydraClient("http://localhost:4445", "http://localhost:4444")
	loginController := controllers.NewLoginController(hydra)

	engine.GET("/login", loginController.Login)
	engine.POST("/login", loginController.AuthUsernamePassword)

	app.engine = engine
}

func (app *App) Run() {
		app.engine.Run(app.port)
}
