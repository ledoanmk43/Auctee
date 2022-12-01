FROM golang:1.18-alpine AS build


WORKDIR /app

COPY go.* /
RUN go mod download

COPY pkg/ /pkg/
COPY src/product-service/ /src/product-service/

RUN CGO_ENABLED=0 GOOS=linux go build -o /main /src/product-service/main.go

FROM scratch

WORKDIR /

COPY --from=build /pkg /pkg/
COPY --from=build /main /main


EXPOSE 1002
EXPOSE 50052

CMD [ "/product-service" ]

ENTRYPOINT ["/main"]