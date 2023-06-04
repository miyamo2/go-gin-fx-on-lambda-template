import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as apiGateway from 'aws-cdk-lib/aws-apigateway';
import * as lambda from 'aws-cdk-lib/aws-lambda';

const API_NAME = "HelloAPI"

interface APIGwHelloAPIResourceProps {
    stage: string;

    helloFunction: lambda.Function
}

export class APIGwHelloAPIResource extends cdk.Resource {
    constructor(scope: cdk.Stack, props: APIGwHelloAPIResourceProps) {
        super(scope, `${API_NAME}Resource-${props.stage}`);
        const api = new apiGateway.RestApi(
            this,
            `HelloAPI-${props.stage}`,
            {
                endpointConfiguration: {
                    types: [apiGateway.EndpointType.REGIONAL]
                },
                deployOptions: {
                    stageName: props.stage,
                }
            }
        )

        // /hello
        const root = api.root.addResource("hello")
        // ANY: ~/hello
        root.addMethod(
            'ANY',
            new apiGateway.LambdaIntegration(props.helloFunction)
        )
        // ANY: ~/hello/{proxy+}
        root.addProxy(
            {
                anyMethod: true,
                defaultIntegration: new apiGateway.LambdaIntegration(props.helloFunction)
            }
        )
    }

}