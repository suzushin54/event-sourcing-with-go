FROM       golang:1.18.3 AS builder

WORKDIR    /go/src/github.com/suzushin54/event-sourcing-with-go
COPY       go.mod go.sum ./

RUN        CGO_ENABLED=0 go build -buildvcs=false -o /bin/app ./internal
ENV        GRPC_HEALTH_PROBE_VERSION v0.4.2
RUN        wget -q -O /bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64
RUN        chmod +x /bin/grpc_health_probe

FROM       golang:1.18.3-alpine3.16 AS deploy
COPY       --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY       --from=builder /app .

CMD ["./app"]

# ---

FROM       golang:1.18.3 AS dev

WORKDIR    /go/src/github.com/suzushin54/event-sourcing-with-go/cmd
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
