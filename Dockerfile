FROM golang:1.15.6

RUN mkdir -p /usr/src/app/

WORKDIR /usr/src/app/

COPY . /usr/src/app/

RUN go build -o FinalGo cmd/web .
FROM golang

ADD . /go/src/

WORKDIR /app

COPY . /app

RUN go build -o app cmd/web/*

EXPOSE 4000

ENTRYPOINT  /app/app
CMD ["./FinalGo"]
