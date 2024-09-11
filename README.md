# Fiap Tech Fast Food

## Project Overview

Fiap Tech Fast Food is a system designed to manage a neighborhood fast food restaurant. The system allows for user registration, product management, order creation, and payment processing. It is built to be resilient to failures and scalable.

This service is responsible for
- Managing the flow of orders in the kitchen
- Controlling the preparation status of dishes
- Coordinating the queue of orders for preparation
- Notifying when orders are ready for pickup
- Monitoring kitchen performance and efficiency

The system is built to be fault resilient and scalable, allowing the kitchen to operate efficiently even during periods of high order volume.

## Badges

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=tech-challenge-fiap-5soat_tc-ff-kitchen-api&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=tech-challenge-fiap-5soat_tc-ff-kitchen-api)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=tech-challenge-fiap-5soat_tc-ff-kitchen-api&metric=bugs)](https://sonarcloud.io/summary/new_code?id=tech-challenge-fiap-5soat_tc-ff-kitchen-api)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=tech-challenge-fiap-5soat_tc-ff-kitchen-api&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=tech-challenge-fiap-5soat_tc-ff-kitchen-api)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=tech-challenge-fiap-5soat_tc-ff-kitchen-api&metric=coverage)](https://sonarcloud.io/summary/new_code?id=tech-challenge-fiap-5soat_tc-ff-kitchen-api)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=tech-challenge-fiap-5soat_tc-ff-kitchen-api&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=tech-challenge-fiap-5soat_tc-ff-kitchen-api)

## Technology Stack

This API was built using [Golang](https://golang.org/) and several tools:
- [gin](http://github.com/gin-gonic/gin) - Web framework
- [mongo-driver](http://go.mongodb.org/mongo-driver) - MongoDB driver
- [viper](https://github.com/spf13/viper) - Configuration tool
- [mockery](https://github.com/vektra/mockery) - Mocking tool for unit tests
- [swag](https://github.com/swaggo/swag) - Swagger documentation generator
- [docker](https://www.docker.com/) - Containerization tool
- [docker-compose](https://docs.docker.com/compose/) - Multi-container Docker applications
- [make](https://www.gnu.org/software/make/) - Task automation tool
- [mermaid](https://mermaid-js.github.io/mermaid/#/) - Diagrams and flowcharts
- [kubernetes](https://kubernetes.io/pt-br/) - Container orchestration

## Architecture

For a demonstration of the architecture, visit: [Architecture Video](AWS_ACCESS_KEY_ID=test;AWS_SECRET_ACCESS_KEY=test)


## Running the Application

### Using Docker

The app can be started using Docker. You can use the pre-defined actions in the Makefile.

#### Build Image

To build an image from the project to push to a registry, use the command below:

```sh
make build-image
```

This command will generate an image with the tag: `tc-ff-kitchen-api`.

#### Generate Documentation

To generate the documentation to publish on the project like an OpenAPI, use the command below:

```sh
make serve-swagger
```

This command will generate a directory called `docs`.

### Development

Before run the application, you need to export the variables below:

```sh
AWS_ACCESS_KEY_ID=test
AWS_SECRET_ACCESS_KEY=test
MONGODB_HOST=localhost
MONGODB_PORT=27017
MONGODB_DATABASE=db
MONGODB_USER=root
MONGODB_PASS=root
```

To run in development for debugging or improvement, use the command:

```sh
make start-local-development &
```

And run that command in another terminal:

```sh
make run
```

This command will start a container with hot-reload for any code modifications, including a container with an instance of MongoDB.

To stop the container, execute:

```sh
make stop-local-development
```

### Testing

Locally, you can use the command below:

```sh
go test ./...  -v
```

Or use a make action:

```sh
make test
```

## Configuration

Configuration settings are managed using environment variables and a configuration file. Refer to the `configs.yaml.sample` file for the required settings.

```yaml:src/external/api/infra/config/configs.yaml.sample
startLine: 1
endLine: 20
```
