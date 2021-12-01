FROM  golang:1.17.3-alpine3.14

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /my-app

CMD ["/my-app"]