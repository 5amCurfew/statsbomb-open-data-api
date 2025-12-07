# Use a small Go base image
FROM golang:1.20-alpine

WORKDIR /app
# Copy and download dependencies first (cache here)
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

COPY . .

RUN go build -o main .
EXPOSE 8080

CMD ["/app/main"]
# docker build -t statsbomb-open-data-api/auth .
# docker run -p 8080:8080 statsbomb-open-data-api/auth