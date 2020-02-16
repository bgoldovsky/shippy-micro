FROM golang:alpine as builder 

RUN apk update && apk upgrade && apk add --no-cache git 

RUN mkdir -p /go/src/app   
WORKDIR /go/src/app    

COPY . .

ENV GO111MODULE=on
RUN go mod init
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippyserver


FROM alpine:latest 

RUN apk --no-cache add ca-certificates

RUN mkdir /app 
WORKDIR /app 
COPY --from=builder /go/src/app/shippyserver .

CMD ["./shippyserver"]