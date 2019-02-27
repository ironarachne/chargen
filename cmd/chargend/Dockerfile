# build stage
FROM golang:1.11 AS build-env
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/chargend

# final stage
FROM scratch
COPY --from=build-env /go/bin/chargend /go/bin/chargend
EXPOSE 9798
CMD ["/go/bin/chargend"]