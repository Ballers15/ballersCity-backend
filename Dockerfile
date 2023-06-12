FROM golang:1.18  as builder

RUN git config \
    --global --add \
    url."https://MKuchum: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@github.com".insteadOf \
    "https://github.com"

RUN apt update
RUN apt install -y protobuf-compiler
RUN go install github.com/luckybet100/protodeps@v1.0.3

WORKDIR /app

COPY ./protodeps-config.yml /app/protodeps-config.yml
RUN protodeps fix
RUN protodeps target build golang

COPY ./go.mod /app/go.mod
RUN go mod download

COPY . /app

RUN make bin/cmd
