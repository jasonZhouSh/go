FROM golang:1.6

RUN go get github.com/labstack/echo/...

RUN mkdir /app
ADD webapi /app/webapi


WORKDIR /app
EXPOSE 1234


