# development stage
FROM golang:1.17-buster AS development
WORKDIR /code
COPY ../go.mod go.sum /code/
RUN go mod download
COPY ../.. /code

# builder stage
FROM development AS builder
RUN go build -o ./app ./cmd/transaction-api

# final stage
FROM gcr.io/distroless/base:nonroot
COPY --from=builder /code/app /
COPY --from=builder /code/db/migrations /db/migrations
ENTRYPOINT [ "/app" ]
