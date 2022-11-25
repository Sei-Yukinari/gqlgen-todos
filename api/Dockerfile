FROM golang:1.19

WORKDIR /app

RUN apt-get update \
    && apt-get install -y wget

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

RUN go install github.com/99designs/gqlgen@v0.17.20 &&  \
    go install github.com/google/wire/cmd/wire@v0.5.0 &&  \
    go install github.com/makiuchi-d/arelo@latest

ENTRYPOINT ["sh", "./scripts/entrypoint.sh"]
