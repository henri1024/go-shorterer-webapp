FROM golang:1.14

WORKDIR /app

EXPOSE 8888

COPY . /app

RUN go mod download

CMD /app/shorterer