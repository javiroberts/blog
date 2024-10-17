# FROM golang:1.23-alpine AS base
FROM golang:1.23-alpine

WORKDIR /app

ADD . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/dist/app .

# FROM gcr.io/distroless/static

# COPY --from=base /app/dist/app .

EXPOSE 1323

CMD [ "/app/dist/app" ]