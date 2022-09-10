FROM golang:1.19-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
# COPY go.mod .
# COPY go.sum .
# RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
# COPY *.go ./

COPY . . 

RUN go mod download

# Build
RUN go build -o ./go-iot-api ./cmd/go-iot-api

# # Run
# dont need this with docker compose
# CMD [ "./go-iot-api" ]
