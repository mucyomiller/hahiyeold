FROM golang:1.10.4-alpine AS builder

WORKDIR /go/src/app
COPY . .

# install git in this alpine
RUN apk update && \
    apk upgrade && \
    apk add git
    
RUN go get -d -v ./...
# RUN go install -v ./...
RUN go build -o hahiyeserver .

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/app/ /app/

EXPOSE 9090

CMD ["./hahiyeserver"]