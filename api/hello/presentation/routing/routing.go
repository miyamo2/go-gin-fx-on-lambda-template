package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/miyamo2theppl/go-gin-fx-on-lambda-template/api/hello/presentation/controller"
)

func createGinEngineFromHelloController(controller *controller.HelloController) *gin.Engine {
	ginengine := gin.Default()
	ginengine.GET("/hello", controller.SayHello)
	ginengine.GET("/hello/:name/greet", controller.DoGreet)
	ginengine.GET("/hello/:name/bye", controller.SayBye)

	return ginengine
}
