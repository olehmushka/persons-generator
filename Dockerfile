# syntax=docker/dockerfile:1
FROM golang:1.26-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /out/persons-generator .

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=builder /out/persons-generator ./persons-generator
COPY config/default.json ./config/default.json
USER nonroot
EXPOSE 8000
ENTRYPOINT ["/app/persons-generator"]
CMD ["http_server_run"]
