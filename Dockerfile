FROM golang:1.20.1-alpine3.17 as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o /finance-api

# FROM alpine:3.17 as publish
FROM scratch

COPY --from=build /finance-api ./

EXPOSE 80

CMD [ "/finance-api","-r" ]