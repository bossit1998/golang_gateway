FROM migrate/migrate:v4.3.1 as builder

FROM golang:1.12-alpine
RUN apk --no-cache add ca-certificates
COPY --from=builder /migrate ${GOPATH}/bin
