FROM golang:1.21
MAINTAINER Christoph Kappestein <christoph.kappestein@apioo.de>
LABEL description="TypeHub Push"

ENV TYPEHUB_CLIENT_ID ""
ENV TYPEHUB_CLIENT_SECRET ""

VOLUME /usr/src/typehub/input

WORKDIR /usr/src/typehub

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/typehub

WORKDIR /usr/src/typehub/input
CMD ["sh", "-c", "/usr/local/bin/typehub pushd /usr/src/typehub/input"]
