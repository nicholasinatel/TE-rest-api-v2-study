# baseimage
FROM golang:1.16 AS builder 

# execute
RUN mkdir /app 
# add all contents from current directory to app directory
ADD . /app 
# new working directory
WORKDIR /app 

# Alreayd have access to Go CLI
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

# Linux baseimage
FROM alpine:latest AS production
# Copy bin compiled file 
COPY --from=builder /app .
# Execute compiled binary
CMD ["./app"]
