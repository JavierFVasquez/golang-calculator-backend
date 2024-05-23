# Golang Test

_Golang Calculator API - Backend application_

## Getting Started 🚀

_In this project, we have a Golang application that deploys a REST API with a serverless architecture and access to a NoSQL database. This ensures scalability, performance, and other software quality attributes (depending on the case)._

## Endpoints (Live 🔴) 🛠️

_These are the endpoints that can be accessed from SwaggerUI, with [this OpenAPI specification](https://raw.githubusercontent.com/JavierFVasquez/golang-calculator-backend/master/open_api_specification.yml):_

### [Swagger Link](https://petstore.swagger.io/?url=https://raw.githubusercontent.com/JavierFVasquez/golang-calculator-backend/master/open_api_specification.yml)

## Prerequisites 📋

- _You need to create a `.env` file in the project root directory with the environment variable for the MongoDB cluster URL and database name._

  ```
  MONGO_URI="mongodb+srv://javierfvasquezr:OPSysYjLRg6drbha@golang.yeghafr.mongodb.net/?retryWrites=true&w=majority"
  DB_NAME="golang"
  ```

- Install [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html) in order to run this project locally
- You need to have Docker installed

  - _Common error (just in case): StoreError('Credentials store docker-credential-desktop exited with "".') Solution: Change ~/.docker/config.json value credsStore to credStore_

- _Another common "error": In local environment yo should to change `libs/auth/auth_middleware.go[18]`_
  ```
  authToken := strings.Replace(request.Headers["authorization"], "Bearer ", "", 1)
  ```
  with
  ```
  authToken := strings.Replace(request.Headers["Authorization"], "Bearer ", "", 1)
  ```

## Installation 🔧

_If you want to perform a step-by-step installation and execution, you should execute the following commands:_

```bash
go mod download (Package installation)
```

```bash
sudo make start-dev (Aplication run)
```

_Pdta: For this step, you need to have SAM CLI and Docker installed and runnning on your machine._

## Testing 🧪

_In order to run unit test, you can run this following command_

```bash
go test ./...   (Run tests)
```

## Built With 🛠️

_These were the tools used for this project:_

- [Serverless](https://www.serverless.com/) - Framework used to create serverless applications on AWS Lambda, providing scalability and cost efficiency.
- [MongoDB](https://www.mongodb.com/) - MongoDB is a flexible, scalable, and NoSQL database that provides high performance and easy data modeling.
- [Cognito](https://aws.amazon.com/es/cognito/) - Cognito is an AWS service for authentication, authorization, and user management, providing secure and scalable user sign-up and sign-in functionality.
- [Mockery](https://github.com/vektra/mockery) - Mockery is a mocking library for Go that simplifies unit testing by creating mock objects of dependencies.
- [Golang](https://go.dev/) -Go is an efficient, concise, and high-performance programming language designed for building scalable and reliable software.

### What I Would Have Liked to Do? ⌨️

_Due to time constraints, there are still things that can be added to this test, and some ideas that come to mind at first glance are:_

- **More Unit Tests**: Unit tests are fundamental to ensure software quality, and strategies like TDD (Test-Driven Development) should be considered depending on the case. They should be implemented with priority. _(Omitted for the test due to time constraints)_
- Security performed with AWS API Gateway authorizers
- Depending on the requirements, decouple models using **Pub-Sub** pattern or **GRPC** in order to comunicate micro-services, or better for example using **Kafka** to process comuinication and queues
- And many more...

## Author ✒️

- **Javier Vasquez** - _Test Author_

---

⌨️ with ❤️ by Javier Vasquez 😊
