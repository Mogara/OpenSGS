FROM golang:1.17 AS builder
ADD . /workspace
WORKDIR /workspace
RUN go mod download
RUN CGO_ENABLED=0  go build -ldflags "-w" -a -o /workspace/openSGS-server cmd/server/main.go

FROM alpine:3.11
WORKDIR /
RUN apk add --no-cache ca-certificates
COPY --from=builder /workspace/openSGS-server /usr/local/bin/openSGS-server
RUN chmod +x /usr/local/bin/openSGS-server
CMD ["sh", "-c", "/usr/local/bin/openSGS-server --host=0.0.0.0 --port=$PORT --log-level=$LOG_LEVEL --allowed-origin=$SHORT_ALLOWED --allowed-origin=$LONG_ALLOWED"]