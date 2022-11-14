FROM golang:1.18-alpine3.16 as builder
LABEL maintainer="wahyudi.ibo.wibowo@gmail.com"
RUN apk add --update --no-cache curl make gcc musl-dev && \
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.0
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./api ./cmd/api/main.go

FROM scratch as app
COPY --from=builder ./app/api ./api
ENTRYPOINT [ "./api" ]