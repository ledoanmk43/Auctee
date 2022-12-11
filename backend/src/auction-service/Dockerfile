#FROM golang:1.18-alpine AS build
#
#ENV GOPATH /auction-service
#
##RUN mkdir /app
##ADD . /app
#WORKDIR /app
#
#COPY go.* /
#RUN go mod download
#
#COPY pkg/ /pkg/
#COPY src/auction-service/ /src/auction-service/
#
#RUN CGO_ENABLED=0 GOOS=linux go build -a -o /main /src/auction-service/main.go
#
#FROM scratch
#
#WORKDIR /
#
#COPY --from=build /pkg /pkg/
#COPY --from=build /main /main
#
#
#EXPOSE 1009
#EXPOSE 50053
#
#ENTRYPOINT ["/main"]

FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.* /
RUN go mod download

COPY pkg/ /pkg/
COPY src/auction-service/ /src/auction-service/
COPY src/account-service/ /src/account-service/
COPY src/product-service/ /src/product-service/

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /main /src/auction-service/main.go

FROM scratch

WORKDIR /

COPY --from=build /main /main

EXPOSE 1009
EXPOSE 50053

ENTRYPOINT ["/main"]