## Base Golang Image for App
FROM golang:1.15

## Create app directory in our image
RUN mkdir /app

## Copy all the directory contents to folder
ADD . /app

## Execute commands from our /app folder
WORKDIR /app

## Build, compile the binary executable
RUN go build -o main .

## Start binary executable
CMD [ "/app/main" ]
