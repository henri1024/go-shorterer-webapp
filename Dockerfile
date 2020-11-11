FROM golang:1.14

WORKDIR /app

EXPOSE 8888

EXPOSE 587

COPY . /app

RUN go mod download

RUN go build

CMD /app/go-shorterer