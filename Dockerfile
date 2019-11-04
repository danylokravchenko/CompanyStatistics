####################################
#   STEP 1 build executable binary  #
#####################################
FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR /app

# Fetch dependencies.
# Using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Fetch dependencies.
# Using go get.
#RUN go get -d -v

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -o main

#####################################
#   STEP 2 build a small image      #
#####################################
FROM scratch

# Copy our static executable.
COPY --from=builder /app/main /app/main

# Copy our config
COPY ./config/config.yaml /config/

EXPOSE 8080

# Run the hello binary.
ENTRYPOINT ["/app/main"]