FROM golang:1.23-alpine as builder

RUN adduser -h /home/iot -D iot

WORKDIR /home/iot

COPY ../../go.mod .
COPY ../../go.sum .
RUN go mod download
COPY ../.. .

RUN chown -R iot:iot /home/iot

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -C ./services/svc_iot/cmd -o /home/iot/iot-manager main.go
RUN chmod 700 /home/iot/iot-manager

## Start a new stage
FROM alpine
RUN adduser -h /home/iot -D iot

WORKDIR /home/iot

COPY --from=builder /home/iot/iot-manager /home/iot/iot-manager
RUN mkdir -p /services/svc_iot/config

COPY ../../services/svc_iot/config/config.yaml ./services/svc_iot/config/config.yaml

RUN chown -R iot:iot /home/iot \
    && chmod -R 700 /services/svc_iot/config

RUN chmod 700 /home/iot/iot-manager

# Run the executable
EXPOSE 4772
CMD ["/home/iot/iot-manager"]