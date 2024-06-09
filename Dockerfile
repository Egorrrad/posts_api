FROM golang:latest

WORKDIR /app

COPY ./ /app



# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client


RUN go mod download
RUN go build -o posts-api ./server


ENTRYPOINT go run ./server