docker build -f ./deploy/Dockerfile --no-cache -t splash-inventory-service .
#docker image prune --filter label=stage=splash-inventory-service-builder