FROM golang:1.26-alpine AS builder
WORKDIR /src
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG VERSION=dev
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w -X main.version=${VERSION}" -o /out/mps-server ./cmd/server

FROM gcr.io/distroless/static-debian12:nonroot
LABEL org.opencontainers.image.title="musician-production-suite"
LABEL org.opencontainers.image.description="Audio production suite backend"
LABEL org.opencontainers.image.licenses="MIT"
WORKDIR /app
COPY --from=builder /out/mps-server /app/mps-server
ENV MPS_API_ADDR=:8080
ENV MPS_STORAGE_DIR=/data/jobs
VOLUME ["/data/jobs"]
EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 CMD ["/app/mps-server", "-healthcheck"]
USER nonroot:nonroot
ENTRYPOINT ["/app/mps-server"]

