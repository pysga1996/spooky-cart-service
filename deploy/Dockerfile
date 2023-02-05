FROM golang:1.20-alpine AS builder
LABEL stage=splash-inventory-service-builder
WORKDIR /tmp/splash-inventory-service
ENV GOOS linux
ENV GOARCH amd64
RUN mkdir "/opt/splash-inventory-service"
RUN mkdir -p "/opt/splash-inventory-service/web/templates"
COPY . /tmp/splash-inventory-service/

RUN go build -gcflags="all=-N -l" -o /opt/splash-inventory-service/splash-inventory-service-v1 github.com/thanh-vt/splash-inventory-service/cmd/splash-inventory-service

FROM golang:1.20-alpine AS main
MAINTAINER thanhvt
WORKDIR /opt/splash-inventory-service
RUN apk add build-base
RUN go install github.com/go-delve/delve/cmd/dlv@latest
COPY --from=builder /opt/splash-inventory-service /opt/splash-inventory-service
ADD ./web /opt/splash-inventory-service/web/
RUN chmod a+x /opt/splash-inventory-service/splash-inventory-service-v1
EXPOSE 8090 9090
CMD dlv --listen=:9090 --headless=true --api-version=2 --accept-multiclient exec /opt/splash-inventory-service/splash-inventory-service-v1