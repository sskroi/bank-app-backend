FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o bank-app cmd/main.go

FROM alpine AS runner

WORKDIR /app

COPY --from=builder /app/bank-app ./
COPY --from=builder /app/configs ./configs/
COPY --from=builder /app/tls ./tls/

CMD ["./bank-app"]
