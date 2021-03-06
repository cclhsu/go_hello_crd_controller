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
3. Open browser on "http://localhost:5000/api/v1/resources" to get list of current resources
4. Use command "docker-compose down" to bring down the docker

#### Run use kubernetes

1. Setup Minikube or kubernetes on docker for desktop
2. Deploy pod and service

Using kubernetes:

```
docker pull cclhsu/go_hello_crd_controller:latest
cd <SRC_ROOT>/go_hello_crd_controller/cmd/crd-controller/k8s
kubectl create -f go-hello-crd-controller-deployment.yaml
kubectl create -f go-hello-crd-controller-service.yaml
```

<!-- ```
# Use only when kube-inject enable automatic
# kubectl create -f go-hello-crd-controller-gateway.yaml
# kubectl create -f go-hello-crd-controller-virtual-service.yaml
``` -->

Using istio:

```
docker pull cclhsu/go_hello_crd_controller:latest
cd <SRC_ROOT>
istioctl kube-inject -f go-hello-crd-controller-istio-spec.yaml | kubectl apply -f -
```
<!-- istioctl kube-inject -f helloworld.yaml | kubectl apply -f - -->

3. Check if deployments and services is working process

Using kubernetes:

```
kubectl describe deployment go-hello-crd-controller
kubectl get deployments
kubectl get pods
kubectl get services
```

Using istio:

```
kubectl get gateways
kubectl get virtualservices
# kubectl get service -n istio-system
kubectl get service istio-ingressgateway -n istio-system
kubectl get service istio-egressgateway -n istio-system
```

4. Check for the exposed port "5000:30549/TCP".  Externally, we will access 30549.

```
15004921clarkhsu:k8s clark_hsu$ kubectl get services
NAME                      TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
go-hello-crd-controller   NodePort    10.98.129.171   <none>        5000:30549/TCP   8s
kubernetes                ClusterIP   10.96.0.1       <none>        443/TCP          34m
```

```
http://localhost:30549/api/v1/resources
```

5. Delete pod and service

Using kubernetes:

```
kubectl delete -f go-hello-crd-controller-service.yaml
kubectl delete -f go-hello-crd-controller-deployment.yaml
```

<!-- ```
kubectl delete -f go-hello-crd-controller-gateway.yaml
kubectl delete -f go-hello-crd-controller-virtual-service.yaml
``` -->

Using istio:

```
kubectl delete --ignore-not-found=true -f go-hello-crd-controller-istio-spec.yaml
```
<!-- kubectl delete -f helloworld.yaml -->

### Use API

#### GET api/v1/resources

Using Docker or K8S:

    curl -i http://localhost:5000/api/v1/resources

Using ISTIO:

    curl -i http://localhost/api/v1/resources

#### GET api/v1/resources/:id

Using Docker or K8S:

    curl -i http://localhost:5000/api/v1/resources/1

Using ISTIO:

    curl -i http://localhost/api/v1/resources/1

#### POST api/v1/resources

Using Docker or K8S:

    curl -i -X POST -H "Content-Type: application/json" -d "{\"name\":\"A\",\"enable\":false,\"state\":\"\"}" http://localhost:5000/api/v1/resources

Using ISTIO:

    curl -i -X POST -H "Content-Type: application/json" -d "{\"name\":\"A\",\"enable\":false,\"state\":\"\"}" http://localhost/api/v1/resources

#### PUT api/v1/resources/:id

Using Docker or K8S:

    curl -i -X PUT -H "Content-Type: application/json" -d "{\"name\":\"B\",\"enable\":false,\"state":\"\"}" http://localhost:5000/api/v1/resources/1

Using ISTIO:

    curl -i -X PUT -H "Content-Type: application/json" -d "{\"name\":\"B\",\"enable\":false,\"state":\"\"}" http://localhost/api/v1/resources/1

#### DELETE api/v1/resources/:id

Using Docker or K8S:

    curl -i -X DELETE http://localhost:5000/api/v1/resources/1

Using ISTIO:

    curl -i -X DELETE http://localhost/api/v1/resources/1

#### GET api/v1/resources/:id/health

Using Docker or K8S:

    curl -i http://localhost:5000/api/v1/resources/1/health

Using ISTIO:

    curl -i http://localhost/api/v1/resources/1/health

#### PUT api/v1/resources/:id/health

Using Docker or K8S:

    curl -i -X PUT -H "Content-Type: application/json" -d "{ \"name\":"A\",\"enable\":true,\"state\":"1\" }" http://localhost:5000/api/v1/resources/1

Using ISTIO:

    curl -i -X PUT -H "Content-Type: application/json" -d "{ \"name\":"A\",\"enable\":true,\"state\":"1\" }" http://localhost/api/v1/resources/1