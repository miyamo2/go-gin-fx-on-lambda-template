import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs'
import { APIGwHelloAPIResource } from './resources/apigw_helloapi_stack'
import { LambdaHelloFunctionResource } from './resources/lambda_hellofunction_stack'

interface GoGinFxOnLambdaTemplateCdkStackProps extends cdk.StackProps {

  rootDir: string;

  stage: string;

}

export class GoGinFxOnLambdaTemplateCdkStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props: GoGinFxOnLambdaTemplateCdkStackProps) {
    super(scope, id, props);

    const lambdaHelloFunctionResource: LambdaHelloFunctionResource = new LambdaHelloFunctionResource(
      this, { "stage": props.stage, "rootDir": props.rootDir }
    )
    const lambdaHelloFunctionResourceOutputs = lambdaHelloFunctionResource.outputValues()
    new APIGwHelloAPIResource(
      this,
      {
        "stage": props.stage,
        "helloFunction": lambdaHelloFunctionResourceOutputs.lambdaFunction
      }
    )
  }
}
