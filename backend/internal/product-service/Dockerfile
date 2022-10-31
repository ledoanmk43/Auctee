FROM golang:1.18.2-alpine AS build

WORKDIR /app

COPY go.* /

RUN go mod download

COPY pkg/ /pkg/

COPY internal/product-service/ /internal/product-service/

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /main /internal/product-service/cmd/main.go

FROM scratch

WORKDIR /

COPY --from=build /pkg /pkg/
COPY --from=build /main /main


EXPOSE 1002
EXPOSE 50052

#CMD [ "/product-service" ]

ENTRYPOINT ["/main"]