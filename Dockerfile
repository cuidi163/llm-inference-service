FROM golang:alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o llm-inference-service

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /app/llm-inference-service .

EXPOSE 8080

CMD ["./llm-inference-service"]