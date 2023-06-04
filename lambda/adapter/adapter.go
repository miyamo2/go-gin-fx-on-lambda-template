package adapter

import (
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

// Create GinAdapter From GinEngine
func FromGinEngine(ginengine *gin.Engine) *ginadapter.GinLambda {
	return ginadapter.New(ginengine)
}
