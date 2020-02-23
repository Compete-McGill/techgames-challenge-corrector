FROM golang:1.13-alpine as builder

RUN mkdir /app 
WORKDIR /app 
COPY . . 
RUN go build cmd/main.go 

FROM node:11.13.0-alpine

RUN apk update && apk upgrade && apk add git
COPY --from=builder /app/main /
ENV API_URL http://localhost:3000
CMD ["/main"]