FROM golang:latest 
RUN mkdir /app
ADD ./app /app/
WORKDIR /app
