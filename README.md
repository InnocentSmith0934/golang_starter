# golang_starter

Based off the Go Writing Web Applications wiki page.

Dockerfiles based off of:
https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/

Static file build command for Dockerfile.small:
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
