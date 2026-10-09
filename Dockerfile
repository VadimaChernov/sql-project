FROM golang:1.25-bookworm

WORKDIR /app

COPY . .

RUN apt-get update && apt-get install -y make

RUN go mod tidy
RUN go build -o /app/exe main.go

CMD ["/app/exe"]