package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/miyamo2theppl/go-gin-fx-on-lambda-template/api/hello/application/usecase"
)

type HelloController struct {
	sayhello usecase.SayHelloUsecase
	dogreet  usecase.DoGreetUsecase
	saybye   usecase.SayByeUsecase
}

func (r *HelloController) SayHello(c *gin.Context) {
	fmt.Println("[DEBUG] START HelloController#SayHello")
	defer fmt.Println("[DEBUG] END HelloController#SayHello")
	usr, err := r.sayhello.SayHello()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usr)
}

func (r *HelloController) DoGreet(c *gin.Context) {
	fmt.Println("[DEBUG] START HelloController#DoGreet")
	defer fmt.Println("[DEBUG] END HelloController#DoGreet")
	name := c.Param("name")
	usr, err := r.dogreet.DoGreet(name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usr)
}

func (r *HelloController) SayBye(c *gin.Context) {
	fmt.Println("[DEBUG] START HelloController#SayBye")
	defer fmt.Println("[DEBUG] END HelloController#SayBye")
	name := c.Param("name")
	usr, err := r.saybye.SayBye(name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, usr)
}

func NewHelloController(
	sayhello usecase.SayHelloUsecase,
	dogreet usecase.DoGreetUsecase,
	saybye usecase.SayByeUsecase,
) *HelloController {
	return &HelloController{
		sayhello: sayhello,
		dogreet:  dogreet,
		saybye:   saybye,
	}
}
