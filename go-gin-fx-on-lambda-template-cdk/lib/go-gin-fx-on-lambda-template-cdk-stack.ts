import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs'
import { APIGwHelloAPIStack } from './resources/apigw_helloapi_stack'
import { LambdaHelloFunctionStack } from './resources/lambda_hellofunction_stack'

interface GoGinFxOnLambdaTemplateCdkStackProps extends cdk.StackProps {

  rootDir: string;

  stage: string;

}

export class GoGinFxOnLambdaTemplateCdkStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props: GoGinFxOnLambdaTemplateCdkStackProps) {
    super(scope, id, props);

    const lambdaHelloFunctionStack: LambdaHelloFunctionStack = new LambdaHelloFunctionStack(
      this, { "stage": props.stage, "rootDir": props.rootDir }
    )
    const lambdaHelloFunctionStackOutputs = lambdaHelloFunctionStack.outputValues()
    new APIGwHelloAPIStack(
      this,
      {
        "stage": props.stage,
        "helloFunction": lambdaHelloFunctionStackOutputs.lambdaFunction
      }
    )
  }
}
