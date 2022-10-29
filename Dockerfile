FROM harbor.front.kjuulh.io/docker-proxy/library/golang:alpine as builder

WORKDIR /src/builder

COPY ci/. .

RUN go build -o dist/dagger-go main.go

FROM harbor.front.kjuulh.io/docker-proxy/library/docker:dind

WORKDIR /src/docker

COPY --from=builder /src/builder/dist/dagger-go .

