FROM golang:latest AS builder

WORKDIR /app

COPY . . 

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /wm-rss cmd/rss/*.go


EXPOSE 8080

CMD ["/wm-rss"]

