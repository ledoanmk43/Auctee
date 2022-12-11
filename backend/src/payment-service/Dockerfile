#FROM golang:1.18-alpine AS build
#
#ENV GOPATH /payment-service
#
##RUN mkdir /app
##ADD . /app
#WORKDIR /app
#
#COPY go.* /
#RUN go mod download
#
#COPY pkg/ /pkg/
#COPY src/payment-service/ /src/payment-service/
#
#RUN CGO_ENABLED=0 GOOS=linux go build -o /main /src/payment-service/main.go
#
#FROM scratch
#
#WORKDIR /
#
#COPY --from=build /pkg /pkg/
#COPY --from=build /main /main
#
#
#EXPOSE 1003
#
##CMD [ "/payment-service" ]
#
#ENTRYPOINT ["/main"]

FROM golang:1.18-alpine AS build

RUN apk add -U --no-cache ca-certificates
WORKDIR /app

COPY go.* /
RUN go mod download

COPY pkg/ /pkg/
COPY src/payment-service/ /src/payment-service/
COPY src/auction-service/ /src/auction-service/
COPY src/account-service/ /src/account-service/
COPY src/product-service/ /src/product-service/

COPY src/payment-service/localhost.pem /src/payment-service/
COPY src/payment-service/localhost-key.pem /src/payment-service/

RUN CGO_ENABLED=0 GOOS=linux go build -o /main /src/payment-service/main.go

FROM scratch

WORKDIR /

COPY --from=build /main /main
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/payment-service/localhost.pem localhost.pem
COPY --from=build /src/payment-service/localhost-key.pem localhost-key.pem

EXPOSE 1003

#CMD [ "/payment-service" ]

ENTRYPOINT ["/main"]