import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as lambda from 'aws-cdk-lib/aws-lambda';
import * as logs from 'aws-cdk-lib/aws-logs';

const FUNCTION_NAME = "HelloFunction"

interface LambdaHelloFunctionResourceProps {
    stage: string;

    rootDir: string;
}

interface LambdaHelloFunctionResourceOutputs {
    lambdaFunction: lambda.Function
}

export class LambdaHelloFunctionResource extends cdk.Resource {

    private readonly outputs: LambdaHelloFunctionResourceOutputs

    constructor(scope: cdk.Stack, props: LambdaHelloFunctionResourceProps) {
        super(scope, `${FUNCTION_NAME}Resource-${props.stage}`);
        const func = new lambda.Function(this, `${FUNCTION_NAME}-${props.stage}`, {
            functionName: `${FUNCTION_NAME}-${props.stage}`,
            runtime: lambda.Runtime.GO_1_X,
            handler: "hello",
            code: lambda.Code.fromAsset(`${props.rootDir}/bin/hello.zip`)
        })

        logs.LogGroup.fromLogGroupName(this, `${FUNCTION_NAME}-${props.stage}LogGroup`, func.logGroup.logGroupName);

        this.outputs = { "lambdaFunction": func }
    }

    /**
     * inspired CFn Outputs Section
     */
    public outputValues(): LambdaHelloFunctionResourceOutputs {
        return this.outputs
    }
}