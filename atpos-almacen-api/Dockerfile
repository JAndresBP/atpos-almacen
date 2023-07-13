FROM golang:1.20

ENV MONGO_SERVER=localhost:27017
ENV RABBITMQ_SERVER=localhost:5672
WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /atpos-almacen-api

EXPOSE 3000

CMD ["/atpos-almacen-api"]