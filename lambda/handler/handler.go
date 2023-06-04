package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

type Handler struct {
	ginLambda *ginadapter.GinLambda
}

func (r Handler) Execute(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return r.ginLambda.ProxyWithContext(ctx, req)
}

func NewHandler(ginLambda *ginadapter.GinLambda) *Handler {
	return &Handler{ginLambda}
}
