FROM golang:1.18.2-alpine AS build

WORKDIR /app

COPY go.* /

RUN go mod download

COPY pkg/ /pkg/

COPY internal/payment-service/ /internal/payment-service/

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /main /internal/payment-service/cmd/main.go

FROM scratch

WORKDIR /

COPY --from=build /pkg /pkg/
COPY --from=build /main /main


EXPOSE 1003
#EXPOSE 50053

#CMD [ "/payment-service" ]

ENTRYPOINT ["/main"]