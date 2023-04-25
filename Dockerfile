##
## Build
##
FROM golang:1.20-alpine as build

# Define working directory
WORKDIR /app

# Install dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy all files
COPY . .

# Build application
RUN CGO_ENABLED=0 go build -o /bin/api cmd/api/main.go

##
## Deploy
##
FROM scratch

ENV APP_HOST=0.0.0.0
ENV APP_PORT=8080
ENV APP_SHUTDOWN_TIMEOUT=10s
ENV GIN_MODE=release

COPY --from=build /bin/api /bin/api

EXPOSE 8080

ENTRYPOINT ["/bin/api"]