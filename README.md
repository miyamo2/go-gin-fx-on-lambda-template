# go-gin-fx-on-lambda-template

Template for building Gin and Fx applications on AWS Lambda.

## Project Structure / プロジェクト構成

```txt
.
├─api
│  └─hello
│      ├─application
│      │  └─usecase
│      ├─configure
│      │  └─fxapp
│      └─presentation
│          ├─controller
│          └─routing
├─fxapp
│  └─runner
│      └─gin
├─lambda
│   ├─adapter
│   ├─handler
│   └─runner
└─go-gin-fx-on-lambda-template-cdk
   ├─bin
   ├─lib
   │  └─resources
   └─test
```

## Class Diagram / クラス図

```mermaid
classDiagram
   class FxAppRunner {
      <<interface>>
      +Run(Lifecycle)* void
   }

   class LambdaRunner {
      -handler *Handler
      +Run(Lifecycle)* void
   }

   class GinRunner {
      -engine *gin.Engine
      -uri    []string
      +Run(Lifecycle)* void
   }

   class Handler {
      -ginLambda *GinLambda
      +Execute(Context, APIGatewayProxyRequest)* APIGatewayProxyResponse, error
   }

   class hello {
      <<LambdaHandler>>
      -app *App
      -main()
   }

   class App {
      <<Fx>>
      +Run()
   }

   class HelloController {
      -sayhello SayHelloUsecase
      -dogreet  DoGreetUsecase
      -saybye   SayByeUsecase
      +SayHello(Context) string
      +DoGreet(Context) string
      +SayBye(Context) string
   }

   class SayHelloUsecase {
      <<interface>>
      +SayHello()* string, error
   }

   class SayHelloUsecaseImpl {
      +SayHello()* string, error
   }

   class DoGreetUsecase {
      <<interface>>
      +DoGreet(string)* string, error
   }

   class DoGreetUsecaseImpl {
      +DoGreet(string)* string, error
   }

   class SayByeUsecase {
      <<interface>>
      +SayBye(string)* string, error
   }

   class SayByeUsecaseImpl {
      +SayBye(string)* string, error
   }

   GinRunner ..|> FxAppRunner
   LambdaRunner ..|> FxAppRunner
   Handler --o LambdaRunner
   SayHelloUsecaseImpl ..|> SayHelloUsecase
   DoGreetUsecaseImpl ..|> DoGreetUsecase
   SayByeUsecaseImpl ..|> SayByeUsecase
   SayHelloUsecase --o HelloController
   DoGreetUsecase --o HelloController
   SayByeUsecase --o HelloController

   App ..> SayHelloUsecaseImpl
   App ..> DoGreetUsecaseImpl
   App ..> SayByeUsecaseImpl
   App ..> LambdaRunner
   App ..> Handler
   App ..> HelloController
   hello ..> App
```
