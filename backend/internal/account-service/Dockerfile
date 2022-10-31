FROM golang:1.18.2-alpine AS build

WORKDIR /app

COPY go.* /

RUN go mod download

COPY pkg/ /pkg/

COPY internal/account-service/ /internal/account-service/

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /main /internal/account-service/cmd/main.go

FROM scratch

WORKDIR /

COPY --from=build /pkg /pkg/
COPY --from=build /main /main


EXPOSE 1001
EXPOSE 50051

#CMD [ "/account-service" ]

ENTRYPOINT ["/main"]