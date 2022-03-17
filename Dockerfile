FROM golang:1.18-alpine

WORKDIR /opt/spooky-cart-service

COPY ./spooky-cart-service-v1 /opt/spooky-cart-service
ADD ./templates /opt/spooky-cart-service/templates
RUN chmod a+x /opt/spooky-cart-service/spooky-cart-service-v1

EXPOSE 8090
CMD [ "/opt/spooky-cart-service/spooky-cart-service-v1" ]