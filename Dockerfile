FROM golang:1.18rc1-alpine3.15 as builder
WORKDIR /app
RUN apk update && apk upgrade && apk add --no-cache ca-certificates git tzdata
RUN update-ca-certificates
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o oncall

FROM scratch
COPY --from=builder /app/oncall .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENTRYPOINT [ "./oncall" ]