Cloud Native Go Application

Microservice created in Go

Overview:
- API simulating requests to retrieve data from books
- Repository mocking data in a slice (no database)
- Docker container created over image 1.10.3-alpine3.8

Docker commands:
- Build container
$ docker build -t cloud-native-go:1.0.0 .
- Run container as a background process
$ docker run --name cloud-native-go -d -p 8080:8080 cloud-native-go:1.0.0


