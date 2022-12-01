FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.* /
RUN go mod download

COPY pkg/ /pkg/
COPY src/payment-service/ /src/payment-service/

RUN CGO_ENABLED=0 GOOS=linux go build -o /main /src/payment-service/main.go

FROM scratch

WORKDIR /

COPY --from=build /pkg /pkg/
COPY --from=build /main /main


EXPOSE 1003

CMD [ "/payment-service" ]

ENTRYPOINT ["/main"]