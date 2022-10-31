FROM harbor.server.kjuulh.io/docker-proxy/library/golang:alpine as builder

WORKDIR /src/builder

COPY tmp/bust .

RUN go build -o dist/bust main.go

FROM harbor.server.kjuulh.io/docker-proxy/library/docker:dind

WORKDIR /src

COPY --from=builder /src/builder/dist/bust /usr/bin/
