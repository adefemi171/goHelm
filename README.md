# Web App



# App Details
A sample appliccation to show how to use Helm with Golang

.

# Proposed Stack

Golang


# Top-level directory layout

    📦go-helm
        📦docs
            ┣ 📜
        📦helm
            ┣ 📜
            ┣ 📜
            ┣ 📜
            ┣ 📜
        📦pkg
            ┣ 📦controller
                ┗ 📜controller.go
            ┣ 📦route
                ┗ 📜route.go
        ┣ 📜.env.example
        ┣ 📜.gitignore
        ┣ 📜docker-compose.yml
        ┣ 📜Dockerfile
        ┣ 📜go.mod
        ┣ 📜go.sum
        ┣ 📜kind.yaml
        ┣ 📜main.go
        ┣ 📜README.md

  


## Folder structure

1. Docs: 

2. helm: This folder contains the helm chart.

4. pkg: This folder contain controllers and routes.



# How to setup project and run locally

### Clone the repository 

```
git clone https://github.com/adefemi171/goHelm.git
```
### Checking Out
The App is built on the ``` main ``` branch you will need to checkout to the app branch using:

```
git checkout main
```

### Run the application

1. Create a .env file with the values below:
```
APP_PORT=:8080
PROMETHEUS_PORT=:9110
```
2. Run the the application using
```
go run main.go
```
3. In your browser, navigate to

```
http://localhost:8080/
```
for the metrics navigate to

```
http://localhost:9110/metrics
```
for the Prometheus metrics


# How to setup project using docker

1. Run the docker-compose.yml file using
```
docker-compose up
```
And visit the endpoint in your browser as specified above

# How to setup project using helm
1. I used [Kind](https://kind.sigs.k8s.io/) to create my cluster
2. In the root folder run:
```
kind create cluster --config kind.yaml
```
3. Change Directory to helm folder and run
```
helm install <application_name> .
```
4. Then watch the pod create with
```
kubectl get pod -w
```
5. Once the pod has been created, you need to ge the service name and port forward to be able to access the container, run
```
kubectl get service
kubectl port-forward svc/<service_name> 8080:8080 9110:9110
```

And visit the endpoint in your browser  as specified above


### Routes

* /header [GET]
* /users [GET]
* /metrics [GET]

#### Example
Visit http://localhost:9110/metrics to see the Prometheus Metric 
## Metric
![](Docs/promMetric.png?raw=true)

or http://localhost:8080/header to see the Response Header
## Response Header
![](Docs/header.png?raw=true)



## Proposed Improvement
1. Set up a pipeline for auto deployment to dockerhub per PR made to main branch

NOTE: [Helm](https://helm.sh/docs/intro/install/), [Docker](https://docs.docker.com/get-docker/), [Golang](https://golang.org/doc/install), [Node](https://nodejs.org/en/download/), and [Kubernetes](https://kubernetes.io/docs/tasks/tools/install-kubectl/) needs to be installed to run this application




