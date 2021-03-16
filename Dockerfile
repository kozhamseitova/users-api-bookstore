FROM golang:1.14

WORKDIR go/scr/Bookstore

COPY . go/src/Bookstore

RUN apt-get update --yes
RUN apt-get install --yes netcat

EXPOSE 8082/tcp