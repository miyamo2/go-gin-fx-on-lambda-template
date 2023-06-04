#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from 'aws-cdk-lib';
import { GoGinFxOnLambdaTemplateCdkStack } from '../lib/go-gin-fx-on-lambda-template-cdk-stack';

const app = new cdk.App();
const rootDir: string = app.node.tryGetContext('rootdir');
const stage: string = app.node.tryGetContext('stage');
new GoGinFxOnLambdaTemplateCdkStack(
  app,
  'GoGinFxOnLambdaTemplateCdkStack',
  {
    rootDir: rootDir,
    stage: stage
  }
);