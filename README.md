# GitHub Actions example using Golang applicaiton

## Jobs created in GitHub actions pipeline

`build` - used to build the whole application
`test` - used to execute the test cases along with checking the test coverage
`deploy` - used to deploy the image of this application to [docker hub](https://hub.docker.com/r/shaikrasheed99/golang-github-actions-example)

CI pipeline defined in this file - [](.github/workflows/pipeline.yaml)

DockerHub link of this application - [](https://hub.docker.com/r/shaikrasheed99/golang-github-actions-example)

## How to run the application

We can run this application by two ways

1. Make sure your machine has golang installed in it. 
    - Run the below command to download the code from Github repository to your local machine.
    ```bash
    git clone https://github.com/shaikrasheed99/golang-github-actions-example.git
    ```
    - Run the below command to move to the application directory.
    ```bash
    cd golang-github-actions-example/
    ``` 
    - Run the below command to install all the dependencies of the application. 
    ```bash
    go mod tidy
    ``` 
    - Run the below command to run the application for testing the APIs. 
    ```bash
    go run .
    ``` 

2. Make sure your machine has docker installed in it. 
    - Run the below command to pull the Image from Dockerhub to your local machine.
    ```bash
    docker pull shaikrasheed99/golang-github-actions-example
    ``` 
    - The Golang application which inside this image is using port 8080.
    - Run the following command in the terminal to start the container.
    ```bash
    docker run -d -p 8080:8080 shaikrasheed99/golang-github-actions-example
    ``` 

## How to test the APIs of application

Totally, we have three simple test APIs

### Test API - 1

#### Request

```
curl --location --request GET 'http://localhost:8080/test1'
```

#### Response

```
"This is first test api"
```

### Test API - 2

#### Request

```
curl --location --request GET 'http://localhost:8080/test2'
```

#### Response

```
"This is second test api"
```

### Test API - 3

#### Request

```
curl --location --request GET 'http://localhost:8080/test3'
```

#### Response

```
"This is third test api"
```