# golang benchmark on Docker (Distroless/Alpine)

Comapres execution time of golang program on docker images with distroless/alpine.

## build

``` bash
# build
docker-compose build

# images
docker images

REPOSITORY                        TAG                 IMAGE ID            CREATED              SIZE
golang-alpine                     latest              e3352ed45dc3        About a minute ago   8.19MB
golang-distroless                 latest              df4d50019df9        About a minute ago   21.8MB
```

## run

```bash
# run on host
go run main.go
     242	   4550611 ns/op

# run on distroless
docker run --rm golang-distroless:latest
     216	   5278074 ns/op

# run on alpine
docker run --rm golang-alpine:latest
     208	   5249024 ns/op
```
