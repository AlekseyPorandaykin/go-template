FROM golang:1.21.1 as build
RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o=/app/bin/application -ldflags="-X 'main.version=$(shell date)'" /app

FROM alpine:latest as app
COPY --from=build /app/bin/application /app/bin/application
COPY --from=build /app/configs /app/configs
RUN chmod +x /app/bin/application

ENTRYPOINT ["/app/bin/application"]