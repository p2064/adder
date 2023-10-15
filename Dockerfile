FROM golang:1.20


ENV GO111MODULE=on

WORKDIR /app
COPY . .

RUN go build -o /app/adder

EXPOSE 9002


# Run
CMD ["/app/adder"]