FROM golang:1.19-alpine

WORKDIR /tmp/splash-inventory-service

ENV GOOS linux
ENV GOARCH amd64

RUN mkdir "/opt/splash-inventory-service"
RUN mkdir -p "/opt/splash-inventory-service/web/templates"

COPY . /tmp/splash-inventory-service/
ADD ./web /opt/splash-inventory-service/web/

RUN go build -o /opt/splash-inventory-service/splash-inventory-service-v1 github.com/thanh-vt/splash-inventory-service/cmd/splash-inventory-service

RUN chmod a+x /opt/splash-inventory-service/splash-inventory-service-v1

EXPOSE 8090
WORKDIR /opt/splash-inventory-service
CMD cd /opt/splash-inventory-service;./splash-inventory-service-v1