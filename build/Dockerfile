FROM golang:1.14.4-alpine3.11

RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached 
# if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the 
# working Directory inside the container 
COPY . .

# Build the Go app
RUN go build -o main

EXPOSE 5000

CMD ["./main"]


