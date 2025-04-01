FROM golang:1.24-alpine AS build

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -ldflags="-w -s" ./...

FROM alpine:latest AS run

WORKDIR /app

COPY --from=build /app/dankstuff.net ./dankstuff.net
COPY --from=build /app/templates ./templates

EXPOSE 8080

CMD ["./dankstuff.net"]
