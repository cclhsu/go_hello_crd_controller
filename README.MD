# CRD Controller

## Instruction

Please help finish below exercise and publish to your github account, update to me when u finish, then I will pull and follow your document to verify. The target date will be next Friday, good luck.

1. Develop a CRD controller to create/update/delete a gin based http service which has /health with actions GET and PUT supported, json payload.
2. By default, a new created resource which will only enable /health action GET.
3. But the CRD has ability to support to update existing resource to enable or disable /heath actions.
4. Unit test is necessary.
5. If have time, take a look istio gateway to think about how to integrate into your CRD to expose your gin http service to outside.

Please provide the readme to describe the design and tutorial about how to execute supported cases.

## Usage

### Run the code

#### Run use local binary

1. Run CRD Controller

```
cd <SRC_ROOT>/go_hello_crd_controller/cmd/crd-controller
go run crud.go
```

2. Run CRD Controller test

```
cd <SRC_ROOT>/go_hello_crd_controller/cmd/crd-controller
go test
```

#### Run use Docker

1. Build Docker
2. Use command "docker-compose up" to bring up the docker
3. Open browser on "http://localhost:8080/api/v1/resources" to get list of current resources
4. Use command "docker-compose down" to bring down the docker

#### Run use kubernetes

TBD

### Use API

#### GET api/v1/resources

    curl -i http://localhost:8080/api/v1/resources

#### GET api/v1/resources/:id

    curl -i http://localhost:8080/api/v1/resources/1

#### POST api/v1/resources

    curl -i -X POST -H "Content-Type: application/json" -d "{\"name\":\"A\",\"enable\":false,\"state\":\"\"}" http://localhost:8080/api/v1/resources

#### PUT api/v1/resources/:id

    curl -i -X PUT -H "Content-Type: application/json" -d "{\"name\":\"B\",\"enable\":false,\"state":\"\"}" http://localhost:8080/api/v1/resources/1

#### DELETE api/v1/resources/:id

    curl -i -X DELETE http://localhost:8080/api/v1/resources/1

#### GET api/v1/resources/:id/health

    curl -i http://localhost:8080/api/v1/resources/1/health

#### PUT api/v1/resources/:id/health

    curl -i -X PUT -H "Content-Type: application/json" -d "{ \"name\":"A\",\"enable\":true,\"state\":"1\" }" http://localhost:8080/api/v1/resources/1