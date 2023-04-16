# syntax=docker/dockerfile:1

FROM golang:1.20
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go install -mod=mod github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon --build="go build main.go" --command="./main"
EXPOSE 8080