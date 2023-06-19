FROM golang

WORKDIR /app
COPY . /app/

RUN go build -o main /app/src/main.go

CMD ["/app/main"]
