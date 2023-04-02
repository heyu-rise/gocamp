FROM golang:alpine
EXPOSE 80
RUN mkdir /app
WORKDIR /app
COPY httpserver /app/httpserver
ENTRYPOINT ["./httpserver"]