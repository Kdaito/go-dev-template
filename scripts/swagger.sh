# !/bin/sh

SWAGGER_FILE="./api/swagger.yaml"

DOCKER_IMAGE="quay.io/goswagger/swagger"

docker run \
  --rm -it --user $(id -u):$(id -g) -v $(pwd):/app -w /app \
  $DOCKER_IMAGE \
  generate model -f $SWAGGER_FILE -m ./internal/application/dto