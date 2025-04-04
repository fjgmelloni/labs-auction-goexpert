FROM golang:1.22.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /auction ./cmd/auction

FROM debian:bullseye-slim

WORKDIR /

COPY --from=builder /auction /auction

# ENV removida porque já está no .env, e docker-compose já injeta
CMD ["/auction"]
