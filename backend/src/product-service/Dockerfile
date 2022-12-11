#FROM golang:1.18-alpine AS build
#
#ENV GOPATH /product-service
#
##RUN mkdir /app
##ADD . /app
#WORKDIR /app
#
#COPY go.* /
#RUN go mod download
#
#COPY pkg/ /pkg/
#COPY src/product-service/ /src/product-service/
#
#RUN CGO_ENABLED=0 GOOS=linux go build -o /main /src/product-service/main.go
#
#FROM scratch
#
#WORKDIR /
#
#COPY --from=build /pkg /pkg/
#COPY --from=build /main /main
#
#
#EXPOSE 1002
#EXPOSE 50052
#
##CMD [ "/product-service" ]
#
#ENTRYPOINT ["/main"]

FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.* /
RUN go mod download

COPY pkg/ /pkg/
COPY src/product-service/ /src/product-service/
COPY src/account-service/ /src/account-service/
COPY src/auction-service/ /src/auction-service/

RUN CGO_ENABLED=0 GOOS=linux go build -o /main /src/product-service/main.go

FROM scratch

WORKDIR /

COPY --from=build /main /main


EXPOSE 1002
EXPOSE 50052

#CMD [ "/product-service" ]

ENTRYPOINT ["/main"]