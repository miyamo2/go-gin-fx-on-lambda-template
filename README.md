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

   class Engine {
      <<Gin>>
   }

   class GinLambda {
      <<ginadapter>>
      -ginEngine *Engine
   }

   class fx_provider {
      <<ResolvedToDI>>
   }

   class invokeOption {
      <<Fx>>
      +Targets []interface
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


   SayHelloUsecaseImpl ..|> SayHelloUsecase
   DoGreetUsecaseImpl ..|> DoGreetUsecase
   SayByeUsecaseImpl ..|> SayByeUsecase
   SayHelloUsecase --o HelloController
   DoGreetUsecase --o HelloController
   SayByeUsecase --o HelloController

   fx_provider ..> SayHelloUsecaseImpl
   fx_provider ..> DoGreetUsecaseImpl
   fx_provider ..> SayByeUsecaseImpl
   fx_provider ..> LambdaRunner
   fx_provider ..> Handler
   fx_provider ..> HelloController
   fx_provider ..> Engine
   fx_provider ..> GinLambda

   Engine ..> HelloController
   GinLambda --o Engine
   LambdaRunner --o Handler
   Handler --o GinLambda

   LambdaRunner ..|> FxAppRunner
   invokeOption ..> FxAppRunner

   App ..> invokeOption
   App ..> fx_provider
   hello ..> App
```
