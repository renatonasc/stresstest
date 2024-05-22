# Executando o build do app
FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o stresstest ./cmd/stresstest

# Gerando a imagem final
FROM alpine:latest
COPY --from=builder /app/stresstest /stresstest
RUN apk add --no-cache ca-certificates
ENTRYPOINT ["/stresstest"]