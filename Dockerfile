# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.23-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /graphql-cocroach-go

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /graphql-cocroach-go /graphql-cocroach-go

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/graphql-cocroach-go"]