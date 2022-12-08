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

WORKDIR /app

COPY go.* /
RUN go mod download

COPY pkg/ /pkg/
COPY src/payment-service/ /src/payment-service/
COPY src/auction-service/ /src/auction-service/
COPY src/account-service/ /src/account-service/
COPY src/product-service/ /src/product-service/

#COPY src/payment-service/minica.pem /src/payment-service/
#COPY src/payment-service/minica-key.pem /src/payment-service/

RUN CGO_ENABLED=0 GOOS=linux go build -o /main /src/payment-service/main.go

FROM scratch

WORKDIR /

COPY --from=build /main /main
#COPY --from=build minica.pem minica.pem
#COPY --from=build minica-key.pem minica-key.pem

EXPOSE 1003

#CMD [ "/payment-service" ]

ENTRYPOINT ["/main"]