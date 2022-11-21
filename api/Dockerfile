FROM golang:1.19

WORKDIR /app

RUN apt-get update \
    && apt-get install -y wget

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

RUN go install github.com/kyoh86/richgo@latest  \
    && go install bitbucket.org/liamstask/goose/cmd/goose@latest \
    && go install github.com/google/wire/cmd/wire@latest \
    && go install github.com/makiuchi-d/arelo@latest

ENTRYPOINT ["sh", "./scripts/entrypoint.sh"]
