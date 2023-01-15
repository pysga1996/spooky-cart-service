FROM golang:1.18-alpine

WORKDIR /opt/splash-inventory-service

COPY ../splash-inventory-service-v1 /opt/splash-inventory-service
ADD ../web/templates /opt/splash-inventory-service/templates
RUN chmod a+x /opt/splash-inventory-service/splash-inventory-service-v1

EXPOSE 8090
CMD [ "/opt/splash-inventory-service/splash-inventory-service-v1" ]