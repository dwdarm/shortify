FROM golang:1.21.7-alpine3.19 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
COPY app ./app

RUN CGO_ENABLED=0 GOOS=linux go build -o docker-shortify

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/docker-shortify ./docker-shortify

ENTRYPOINT ["./docker-shortify"]