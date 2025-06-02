
---

### `Dockerfile`
```dockerfile
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o bridge

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bridge .
EXPOSE 8080

# Konfigurierbare Umgebung
ENV HOMEMATIC_URL=""
ENV HOMEMATIC_TOKEN=""

CMD ["./bridge"]
