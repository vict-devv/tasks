FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN make build

FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /app/bin/tasks-api .
COPY --from=builder /app/config.yaml .

ENV PATH=$PATH:/app

ARG SERVER_PORT=8080
EXPOSE ${SERVER_PORT}

ENTRYPOINT [ "tasks-api" ]