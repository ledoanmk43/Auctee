#FROM golang:1.18-alpine AS build
#
#ENV GOPATH /account-service
#
##RUN mkdir /app
##ADD . /app
#WORKDIR /app
#
#COPY go.* /
#RUN go mod download
#
#COPY pkg/ /pkg/
#COPY src/account-service/ /src/account-service/
#
#RUN CGO_ENABLED=0 GOOS=linux go build -o /main /src/account-service/main.go
#
#FROM scratch
#
#WORKDIR /
#
#COPY --from=build /pkg /pkg/
#COPY --from=build /main /main
#
#
#EXPOSE 1001
#EXPOSE 50051
#
##CMD [ "/account-service" ]
#
#ENTRYPOINT ["/main"]

FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.* /
RUN go mod download

COPY pkg/ /pkg/
COPY src/account-service/ /src/account-service/

RUN CGO_ENABLED=0 GOOS=linux go build -o /main /src/account-service/main.go

FROM scratch

WORKDIR /

COPY --from=build /main /main


EXPOSE 1001
EXPOSE 50051

#CMD [ "/account-service" ]

ENTRYPOINT ["/main"]