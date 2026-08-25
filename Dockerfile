FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY shared/go shared/go
COPY backend/robot-service backend/robot-service
WORKDIR /app/backend/robot-service
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/backend/robot-service/main .
COPY backend/robot-service/database/migrations/ ./database/migrations/
CMD ["./main"]
