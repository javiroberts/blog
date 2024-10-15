FROM golang:1.23-alpine

WORKDIR /app

ADD . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/blog

CMD [ "./blog" ]