FROM golang:1.16

WORKDIR /app

COPY . .

RUN go env -w GO111MODULE=off && go build -o math

CMD ["./math"]
