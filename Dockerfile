FROM golang:1.17-alpine

RUN mkdir /check-location
WORKDIR /check-location

COPY go.mod .

COPY go.sum .

RUN go mod download

ADD . /check-location/

RUN chmod -R 777 /check-location

RUN adduser -S -D -H -h /check-location appuser

USER appuser

RUN go build -o checkLocation .

EXPOSE 3005

CMD ["./checkLocation"]