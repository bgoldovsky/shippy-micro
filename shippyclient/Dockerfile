FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir -p /go/src/app   
WORKDIR /go/src/app

COPY . .

ENV GO111MODULE=on
RUN go mod init
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippyclient


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
ADD consignment.json /app/consignment.json
COPY --from=builder /go/src/app/shippyclient .

CMD ["./shippyclient"]