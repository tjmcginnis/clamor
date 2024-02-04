FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./
ADD cmd ./cmd
ADD handlers ./handlers
ADD templates ./templates

RUN CGO_ENABLED=0 GOOS=linux go build -o /clamor ./cmd/clamor

CMD ["/clamor"]
