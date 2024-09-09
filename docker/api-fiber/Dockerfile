FROM golang:1.22.6-alpine

WORKDIR /api-fiber

ARG ENVIRONMENT

# Copy package files and download
COPY api-fiber/go.mod api-fiber/go.sum ./
RUN go mod download

# Copy source code
COPY api-fiber/config ./config
COPY api-fiber/controllers ./controllers
COPY api-fiber/libs ./libs
COPY api-fiber/middlewares ./middlewares
COPY api-fiber/models ./models
COPY api-fiber/routes ./routes
COPY api-fiber/services ./services
COPY api-fiber/main.go ./

# Build and run
RUN go build .
ENTRYPOINT ["./api-fiber"]
