FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./
ADD templates ./templates

RUN CGO_ENABLED=0 GOOS=linux go build -o /clamor

CMD ["/clamor"]
