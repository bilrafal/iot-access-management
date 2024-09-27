FROM golang:1.23-alpine as builder

RUN adduser -h /home/cred-admin -D cred-admin

WORKDIR /home/cred-admin

COPY ../../go.mod .
COPY ../../go.sum .
RUN go mod download
COPY ../.. .

RUN chown -R cred-admin:cred-admin /home/cred-admin

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -C ./services/svc_credential_manager/cmd -o /home/cred-admin/credential-manager main.go
RUN chmod 700 /home/cred-admin/credential-manager

## Start a new stage
FROM alpine
RUN adduser -h /home/cred-admin -D cred-admin

WORKDIR /home/cred-admin

COPY --from=builder /home/cred-admin/credential-manager /home/cred-admin/credential-manager
RUN mkdir -p /services/svc_credential_manager/config

COPY ../../services/svc_credential_manager/config/config.yaml ./services/svc_credential_manager/config/config.yaml

RUN chown -R cred-admin:cred-admin /home/cred-admin \
    && chmod -R 700 /services/svc_credential_manager/config

RUN chmod 700 /home/cred-admin/credential-manager

# Run the executable
EXPOSE 4771
CMD ["/home/cred-admin/credential-manager"]