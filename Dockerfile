FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -o ./dorker .

FROM alpine
WORKDIR /app
COPY .env .env
COPY --from=builder /app/dorker ./dorker
CMD ./dorker