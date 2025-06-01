
---

### `Dockerfile`
```dockerfile
FROM golang:1.22-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o bridge

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/bridge .
ENV HOMEMATIC_URL=""
ENV HOMEMATIC_TOKEN=""
EXPOSE 8080
CMD ["./bridge"]
