FROM golang:1.20

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /cordbleibaum-backend

EXPOSE 3000

CMD ["/cordbleibaum-backend"]
